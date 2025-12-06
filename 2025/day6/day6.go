package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Op string

const (
	Plus Op = "+"
	Mult    = "*"
)

type Problem struct {
	nums []int
	op   Op
}

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

	problems := parsePart1(f)
	res1 := calcRes(problems)
	fmt.Printf("Part1: %d\n", res1)

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	problems = parsePart2(string(data))
	res2 := calcRes(problems)
	fmt.Printf("Part2: %d\n", res2)
}

func parsePart2(content string) []Problem {
	lines := strings.Split(content, "\n")
	maxLen := 0
	for _, line := range lines {
		maxLen = max(maxLen, len(line))
	}

	problems := make([]Problem, 0, len(lines[0])) // to big, but fine
	i := 0
	for col := 0; col < maxLen; col++ {
		var numStr string
		for _, line := range lines[:len(lines)-2] {
			if len(line) <= col {
				break
			}
			numStr += string(line[col])
		}
		numStr = strings.Trim(numStr, " ")
		if numStr == "" {
			i += 1
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		if len(problems) == i {
			problems = append(problems, Problem{})
		}
		problems[i].nums = append(problems[i].nums, num)
	}
	for j, op := range strings.Fields(lines[len(lines)-2]) {
		problems[j].op = Op(op)
	}
	return problems
}

func calcRes(problems []Problem) int {
	var res int
	for _, problem := range problems {
		res += problem.Do()
	}
	return res
}

func parsePart1(f io.Reader) []Problem {
	var problems []Problem

	scanner := bufio.NewScanner(f)
	start := true
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if start {
			problems = make([]Problem, len(parts))
			start = false
		}
		if len(parts) != len(problems) {
			panic("length mismtach")
		}
		if parts[0] == string(Plus) || parts[0] == string(Mult) {
			for i := 0; i < len(parts); i++ {
				op := Plus
				if parts[i] == string(Mult) {
					op = Mult
				}
				problems[i].op = op
			}
		} else {
			for i := 0; i < len(parts); i++ {
				num, err := strconv.Atoi(parts[i])
				if err != nil {
					log.Fatal(err)
				}
				problems[i].nums = append(problems[i].nums, num)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return problems
}

func (p Problem) Do() int {
	var res int
	for _, num := range p.nums {
		if p.op == Plus {
			res += num
		} else if p.op == Mult {
			if res == 0 {
				res++
			}
			res *= num
		} else {
			panic("unknown op")
		}
	}
	return res
}
