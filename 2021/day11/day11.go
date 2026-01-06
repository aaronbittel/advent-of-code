package main

import (
	"AOC2021/internal/common"
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	SIZE      = 10
	MAX_LEVEL = 10
)

type Octopus struct {
	Level   int
	Flashed bool
}

type Grid [SIZE][SIZE]*Octopus

type Point struct {
	Y int
	X int
}

var DIRECTIONS = []Point{
	{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1},
}

func main() {
	f := common.GetFile()
	defer f.Close()

	grid := parse(f)

	part1Grid := copyArray(grid)
	part1, dur1 := common.TimeIt(func() int {
		var res int
		for range 100 {
			res += flashes(&part1Grid)
		}
		return res
	})
	fmt.Printf("Part1: %d, took %s\n", part1, dur1)

	part2, dur2 := common.TimeIt(func() int {
		var (
			res int
			r   int
		)
		for ; r != 100; res++ {
			r = flashes(&grid)
			fmt.Printf("\r%d %% flashing.", r)
		}
		fmt.Printf("\r\033[2K")
		return res
	})
	fmt.Printf("Part2: %d, took %s\n", part2, dur2)
}

func flashes(grid *Grid) int {
	var res int

	// Step 1
	for y := range SIZE {
		for x := range SIZE {
			grid[y][x].Level++
		}
	}

	// Step 2
	for y := range SIZE {
		for x := range SIZE {
			octo := grid[y][x]
			if octo.Level < MAX_LEVEL {
				continue
			}
			if !octo.Flashed {
				flash(grid, y, x)
			}
		}
	}

	// Step 3
	for y := range SIZE {
		for x := range SIZE {
			octo := grid[y][x]
			if octo.Level >= MAX_LEVEL {
				res++
				octo.Level = 0
			}
			octo.Flashed = false
		}
	}

	return res
}

func flash(grid *Grid, y, x int) {
	grid[y][x].Flashed = true

	for _, dir := range DIRECTIONS {
		newY, newX := y+dir.Y, x+dir.X
		if !grid.InBounds(newY, newX) {
			continue
		}
		octo := grid[newY][newX]
		octo.Level++
		if octo.Level >= MAX_LEVEL && !octo.Flashed {
			octo.Flashed = true
			flash(grid, newY, newX)
		}
	}
}

func parse(r io.Reader) Grid {
	grid := Grid{}

	scanner := bufio.NewScanner(r)
	for y := 0; scanner.Scan(); y++ {
		for x, ch := range scanner.Text() {
			grid[y][x] = &Octopus{Level: int(ch - '0')}
		}
	}

	return grid
}

const (
	LIGHTGRAY = "\033[90m"
	RESET     = "\033[0m"
)

func (g Grid) String() string {
	sb := strings.Builder{}

	for y := range SIZE {
		for x := range SIZE {
			octo := g[y][x]
			if octo.Level != 0 {
				sb.WriteString(LIGHTGRAY)
			}
			sb.WriteByte(byte('0' + octo.Level))
			if octo.Level != 0 {
				sb.WriteString(RESET)
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (g Grid) InBounds(y, x int) bool {
	return y >= 0 && y < SIZE && x >= 0 && x < SIZE
}

func copyArray(src Grid) Grid {
	dst := Grid{}
	for y := range SIZE {
		for x := range SIZE {
			o := src[y][x]
			dst[y][x] = &Octopus{Level: o.Level}
		}
	}
	return dst
}
