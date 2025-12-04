package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Grid struct {
	data   []bool
	height int
	width  int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <input-file>\n", os.Args[0])
		os.Exit(1)
	}

	parseStart := time.Now()

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	content := string(data)
	grid := parse(content)
	parseDur := time.Since(parseStart)

	part1Start := time.Now()
	res1 := part1(grid)
	part1Dur := time.Since(part1Start)
	part2Start := time.Now()
	res2 := part2(grid)
	part2Dur := time.Since(part2Start)

	fmt.Printf("Parsing took: %s\n", parseDur)
	fmt.Printf("Part1: %d, took: %s\n", res1, part1Dur)
	fmt.Printf("Part2: %d, took: %s\n", res2, part2Dur)
}

func part2(grid Grid) int {
	var result int
	for {
		var toBeRemoved []int
		for y := range grid.height {
			for x := range grid.width {
				if !grid.At(y, x) {
					continue
				}
				if grid.CountNeighbours(y, x) < 4 {
					result++
					toBeRemoved = append(toBeRemoved, y*grid.width+x)
				}
			}
		}

		if len(toBeRemoved) == 0 {
			return result
		}

		for _, i := range toBeRemoved {
			grid.data[i] = false
		}
	}
}

func part1(grid Grid) int {
	var result int
	for y := range grid.height {
		for x := range grid.width {
			if !grid.At(y, x) {
				continue
			}
			if grid.CountNeighbours(y, x) < 4 {
				result++
			}
		}
	}
	return result
}

func parse(content string) Grid {
	lines := strings.Split(content, "\n")
	height, width := len(lines)-1, len(lines[0])
	data := make([]bool, height*width)

	for y, line := range lines {
		for x, char := range line {
			data[y*width+x] = char != '.'
		}
	}

	return Grid{
		data:   data,
		height: height,
		width:  width,
	}
}

func (g Grid) CountNeighbours(y, x int) int {
	var count int
	for dy := -1; dy < 2; dy++ {
		for dx := -1; dx < 2; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			if g.At(y+dy, x+dx) {
				count++
			}
		}
	}
	return count
}

func (g Grid) At(y, x int) bool {
	if y < 0 || y >= g.height || x < 0 || x >= g.width {
		return false
	}
	return g.data[y*g.width+x]
}

func (g Grid) String() string {
	b := strings.Builder{}
	for y := range g.height {
		for x := range g.width {
			char := "."
			if g.data[y*g.width+x] {
				char = "@"
			}
			fmt.Fprintf(&b, char)
		}
		fmt.Fprintln(&b)
	}
	return b.String()
}
