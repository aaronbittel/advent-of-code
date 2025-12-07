package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <input-file>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	res1 := part1(f)
	fmt.Printf("Part1: %d\n", res1)
	f.Seek(0, io.SeekStart)

	res2 := part2(f)
	fmt.Printf("Part2: %d\n", res2)
}

func part2(f io.Reader) int {
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	first := scanner.Text()
	beams := map[int]int{strings.Index(first, "S"): 1}
	for scanner.Scan() {
		line := scanner.Text()
		for idx := range beams {
			if line[idx] != '^' {
				continue
			}
			beams[idx+1] += beams[idx]
			beams[idx-1] += beams[idx]
			beams[idx] = 0
		}
	}
	var res int
	for _, v := range beams {
		res += v
	}
	return res
}

func part1(f io.Reader) int {
	var res int

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	beams := map[int]struct{}{
		strings.Index(line, "S"): {},
	}
	for scanner.Scan() {
		line := scanner.Text()
		newBeams := make([]int, 0, len(beams)*2)
		for idx := range beams {
			if line[idx] != '^' {
				continue
			}
			res++
			delete(beams, idx)
			newBeams = append(newBeams, idx-1, idx+1)
		}
		for _, b := range newBeams {
			beams[b] = struct{}{}
		}
	}
	return res
}
