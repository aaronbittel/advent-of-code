package main

import (
	"AOC2022/internal/common"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Grid [][]byte

type Point struct {
	Y, X int
}

var DIRECTIONS = []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func main() {
	filename := common.GetFilename()
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	grid, start, goal := parse(data)

	res1, dur1 := common.TimeIt(func() int {
		return shortestLength(grid, start, goal)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)

	res2, dur2 := common.TimeIt(func() int {
		return part2(grid, goal)
	})
	fmt.Printf("Part2: %d, took %s\n", res2, dur2)
}

func part2(grid Grid, goal Point) int {
	res := math.MaxInt
	for y, row := range grid {
		for x, ch := range row {
			if ch != 'a' {
				continue
			}
			start := Point{Y: y, X: x}
			length := shortestLength(grid, start, goal)
			if length != -1 {
				res = min(res, length)
			}
		}
	}
	return res
}

func shortestLength(grid Grid, start, goal Point) int {
	res := -1

	type State struct {
		pos    Point
		length int
	}
	queue := []State{{pos: start}}
	visited := make(map[Point]struct{})
	visited[start] = struct{}{}

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		if s.pos == goal {
			res = s.length
			break
		}

		for _, dir := range DIRECTIONS {
			newY, newX := s.pos.Y+dir.Y, s.pos.X+dir.X
			if newY < 0 || newY >= len(grid) || newX < 0 || newX >= len(grid[0]) {
				continue
			}
			if grid[newY][newX] > grid[s.pos.Y][s.pos.X]+1 {
				continue
			}
			newPos := Point{Y: newY, X: newX}
			if _, ok := visited[newPos]; ok {
				continue
			}
			visited[newPos] = struct{}{}
			queue = append(queue, State{pos: newPos, length: s.length + 1})
		}
	}
	return res
}

func parse(data []byte) (Grid, Point, Point) {
	data = bytes.TrimRight(data, "\n")
	lines := bytes.Split(data, []byte{'\n'})

	var start, goal Point

	grid := make([][]byte, len(lines))
	for y, line := range lines {
		grid[y] = make([]byte, len(line))
		for x, ch := range line {
			grid[y][x] = ch
			if ch == 'S' {
				start = Point{Y: y, X: x}
				grid[y][x] = 'a'
			}
			if ch == 'E' {
				goal = Point{Y: y, X: x}
				grid[y][x] = 'z'
			}
		}
	}
	return grid, start, goal
}

func (g Grid) String() string {
	sb := strings.Builder{}
	sb.Grow(len(g) * len(g[0]))

	for _, row := range g {
		for _, ch := range row {
			sb.WriteByte(ch)
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}
