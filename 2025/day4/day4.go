package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Point struct {
	x int
	y int
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

func part2(grid map[Point]rune) int {
	var result int
	var toBeRemoved []Point
	for p, s := range grid {
		if s == '.' {
			continue
		}
		if countNeighbours(grid, p) < 4 {
			result++
			toBeRemoved = append(toBeRemoved, p)
		}
	}

	if len(toBeRemoved) == 0 {
		return result
	}

	for _, p := range toBeRemoved {
		grid[p] = '.'
	}

	return result + part2(grid)
}

func part1(grid map[Point]rune) int {
	var result int
	for p, s := range grid {
		if s == '.' {
			continue
		}
		if countNeighbours(grid, p) < 4 {
			result++
		}
	}
	return result
}

func countNeighbours(grid map[Point]rune, p Point) int {
	var count int
	for y := -1; y < 2; y++ {
		for x := -1; x < 2; x++ {
			if y == 0 && x == 0 {
				continue
			}
			new_p := Point{x: p.x + x, y: p.y + y}
			if sym, ok := grid[new_p]; ok && sym == '@' {
				count += 1
			}
		}
	}
	return count
}

func parse(content string) map[Point]rune {
	lines := strings.Split(content, "\n")
	grid := make(map[Point]rune, len(lines)*len(lines[0]))

	for y, line := range lines {
		for x, char := range line {
			grid[Point{x: x, y: y}] = char
		}
	}

	return grid
}
