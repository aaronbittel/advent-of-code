package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	patterns, designs := parse(f)

	res1, dur1 := timeIt(func() int {
		return part1(patterns, designs)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)

	res2, dur2 := timeIt(func() int {
		return part2(patterns, designs)
	})
	fmt.Printf("Part2: %d, took %s\n", res2, dur2)
}

func timeIt[T any](f func() T) (T, time.Duration) {
	start := time.Now()
	res := f()
	dur := time.Since(start)
	return res, dur
}

func part2(patterns, designs []string) int {
	var res int
	for _, design := range designs {
		possPatterns := calcPossPatterns(patterns, design)
		if checkDesign(possPatterns, design) {
			res += allDesigns(possPatterns, design)
		}
	}
	return res
}

func calcPossPatterns(patterns []string, design string) []string {
	poss := make([]string, 0, len(patterns))
	for _, p := range patterns {
		if strings.Contains(design, p) {
			poss = append(poss, p)
		}
	}
	return poss
}

func allDesigns(patterns []string, design string) int {
	memo := make(map[string]int)

	var dfs func(string) int
	dfs = func(d string) int {
		if d == "" {
			return 1
		}

		if v, ok := memo[d]; ok {
			return v
		}

		var count int
		for _, p := range patterns {
			if strings.HasPrefix(d, p) {
				count += dfs(d[len(p):])
			}
		}

		memo[d] = count
		return count
	}

	return dfs(design)
}

func part1(patterns, designs []string) int {
	var res int
	for _, design := range designs {
		if checkDesign(patterns, design) {
			res++
		}
	}
	return res
}

func checkDesign(patterns []string, design string) bool {
	memo := make(map[string]bool)

	dfs := func(design string) bool {
		if design == "" {
			return true
		}
		for _, pattern := range patterns {
			if v, ok := memo[design]; ok {
				return v
			}
			if strings.HasPrefix(design, pattern) {
				if checkDesign(patterns, design[len(pattern):]) {
					memo[design] = true
					return true
				}
			}
		}
		memo[design] = false
		return false
	}

	return dfs(design)
}

func parse(f io.Reader) ([]string, []string) {
	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		panic("wrong input")
	}
	patterns := strings.Split(scanner.Text(), ", ")
	if !scanner.Scan() {
		panic("expected empty line")
	}

	designs := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		designs = append(designs, line)
	}

	return patterns, designs
}
