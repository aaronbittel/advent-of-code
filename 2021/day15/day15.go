package main

import (
	"AOC2021/internal/common"
	"bytes"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Grid [][]int

func main() {
	filename := common.GetFilename()
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	grid := parse(data)

	res1, dur1 := common.TimeIt(func() int {
		return part1(grid)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)
}

type Point struct {
	Y int
	X int
}

var DIRECTIONS = []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func part1(grid Grid) int {
	goal := Point{Y: len(grid) - 1, X: len(grid[0]) - 1}
	var res int

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &State{})
	memo := make(map[Point]int)
	memo[Point{}] = 0

	for pq.Len() > 0 {
		s := heap.Pop(pq).(*State)

		if s.P == goal {
			res = s.Risk
			break
		}

		for _, dir := range DIRECTIONS {
			newY, newX := s.P.Y+dir.Y, s.P.X+dir.X
			if !grid.InBounds(newY, newX) {
				continue
			}
			newP := Point{Y: newY, X: newX}
			newRisk := s.Risk + grid[newY][newX]
			if r, ok := memo[newP]; ok {
				if r <= newRisk {
					continue
				}
			}
			memo[newP] = newRisk
			heap.Push(pq, &State{P: newP, Risk: newRisk})
		}
	}

	return res
}

func parse(data []byte) Grid {
	lines := bytes.Split(data, []byte("\n"))
	lines = lines[:len(lines)-1] // cut last empty line
	grid := make([][]int, len(lines))

	for y, line := range lines {
		row := make([]int, len(line))
		for x, b := range line {
			row[x] = int(b - '0')
		}
		grid[y] = row
	}

	return grid
}

func (g Grid) String() string {
	sb := strings.Builder{}
	for _, row := range g {
		for _, v := range row {
			sb.WriteString(strconv.Itoa(v))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (g Grid) InBounds(y, x int) bool {
	return y >= 0 && y < len(g) && x >= 0 && x < len(g[0])
}
