package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	// Sort by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	// Merge and accumulate
	mergedStart := ranges[0].start
	mergedEnd := ranges[0].end
	total := 0

	for _, r := range ranges[1:] {
		if r.start > mergedEnd {
			// no overlap -> close current interval
			total += mergedEnd - mergedStart + 1

			// start a new interval
			mergedStart = r.start
			mergedEnd = r.end
		} else {
			// overlap -> extend the end if needed
			if r.end > mergedEnd {
				mergedEnd = r.end
			}
		}
	}

	// close final interval
	total += mergedEnd - mergedStart + 1
	return total
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
