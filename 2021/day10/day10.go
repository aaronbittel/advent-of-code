package main

import (
	"AOC2021/internal/common"
	"bufio"
	"fmt"
	"slices"
	"time"
)

var (
	OPEN_TAGS    = []byte{'(', '[', '{', '<'}
	CLOSE_TAGS   = []byte{')', ']', '}', '>'}
	ERROR_POINTS = map[byte]int{
		')': 3, ']': 57, '}': 1197, '>': 25137,
	}
	COMPLETION_POINTS = map[byte]int{
		')': 1, ']': 2, '}': 3, '>': 4,
	}
)

func main() {
	start := time.Now()
	f := common.GetFile()
	defer f.Close()

	var (
		part1       int
		part2Scores = make([]int, 0, 128)
	)

	scanner := bufio.NewScanner(f)
outer:
	for scanner.Scan() {
		var p2Score int
		line := scanner.Text()
		buf := make([]byte, 0, len(line))
		for i := range line {
			b := line[i]
			if slices.Contains(OPEN_TAGS, b) {
				buf = append(buf, b)
			} else {
				if len(buf) == 0 || !matchTag(pop(&buf), b) {
					part1 += ERROR_POINTS[b]
					continue outer
				}
			}
		}
		for _, b := range slices.Backward(buf) {
			p2Score = p2Score*5 + COMPLETION_POINTS[findMatch(b)]
		}
		part2Scores = append(part2Scores, p2Score)
	}

	slices.Sort(part2Scores)
	part2 := part2Scores[len(part2Scores)/2]

	dur := time.Since(start)

	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
	fmt.Printf("Part1 & Part2 took %s\n", dur)
}

func findMatch(openTag byte) byte {
	if openTag == '(' {
		return ')'
	}
	if openTag == '[' {
		return ']'
	}
	if openTag == '{' {
		return '}'
	}
	if openTag == '<' {
		return '>'
	}

	panic("illegal opentag")
}

func matchTag(openTag, closeTag byte) bool {
	if openTag == '(' && closeTag == ')' {
		return true
	}
	if openTag == '[' && closeTag == ']' {
		return true
	}
	if openTag == '{' && closeTag == '}' {
		return true
	}
	if openTag == '<' && closeTag == '>' {
		return true
	}
	return false
}

func pop(bs *[]byte) byte {
	last := len((*bs)) - 1
	b := (*bs)[last]
	(*bs) = (*bs)[:last]
	return b
}
