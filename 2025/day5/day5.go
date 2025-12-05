package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <input-file>\n", os.Args[0])
		os.Exit(1)
	}

	ranges, ids := parse(os.Args[1])
	res1 := part1(ranges, ids)
	fmt.Printf("Part1: %d\n", res1)

	res2 := part2(ranges)
	fmt.Printf("Part2: %d\n", res2)
}

func part2(ranges []Range) int {
	var res int

	for len(ranges) > 0 {
		smallestIdx := 0
		imRange := ranges[smallestIdx]
		for i, r := range ranges {
			if r.start < imRange.start {
				smallestIdx = i
				imRange = r
			}
		}

		marked := make(map[int]struct{})
		marked[smallestIdx] = struct{}{}
		// combine all ranges that have overlap with imRange
		for {
			changed := false
			for i, r := range ranges {
				// completly distinct
				if r.start > imRange.end {
					continue
				}

				marked[i] = struct{}{}
				// already covered by imRange
				if r.start >= imRange.start && r.end <= imRange.end {
					continue
				}
				imRange.end = r.end
				changed = true
				break
			}

			if !changed {
				break
			}
		}

		res += imRange.end - imRange.start + 1

		cleaned := []Range{}
	outer:
		for i, r := range ranges {
			for m := range marked {
				if i == m {
					continue outer
				}
			}
			cleaned = append(cleaned, r)
		}

		ranges = cleaned
	}

	return res
}

func part1(ranges []Range, ids []int) int {
	var res int

	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.start && id <= r.end {
				res++
				break
			}
		}
	}

	return res
}

func parse(filename string) ([]Range, []int) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var ranges []Range
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		ranges = append(ranges, Range{start: start, end: end})
	}

	var ids []int
	for scanner.Scan() {
		line := scanner.Text()
		id, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		ids = append(ids, id)
	}

	return ranges, ids
}
