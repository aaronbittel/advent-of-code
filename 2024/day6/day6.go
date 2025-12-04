package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Dir int

const (
	Up    Dir = 0b0001
	Right     = 0b0010
	Down      = 0b0100
	Left      = 0b1000
)

type Guard struct {
	y      int
	x      int
	facing Dir
}

type Grid struct {
	data   []rune
	height int
	width  int
	guard  Guard
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	content := string(data)
	parseStart := time.Now()
	grid := parse(content)
	parseDur := time.Since(parseStart)
	fmt.Printf("Parsing took %s\n", parseDur)

	part1Grid := grid.Copy()
	part1Start := time.Now()
	part1(part1Grid)
	res1 := part1Grid.Count()
	part1Dur := time.Since(part1Start)
	fmt.Printf("Part1: %d, took %s\n", res1, part1Dur)

	part2Start := time.Now()
	part2Grid := grid.Copy()
	part1(part2Grid)
	path := part2Grid.Path()
	res2 := part2(grid, path)
	part2Dur := time.Since(part2Start)
	fmt.Printf("\r%s", strings.Repeat(" ", 20))
	fmt.Printf("\rPart2: %d, took %s\n", res2, part2Dur)

}

func part2(grid Grid, path []int) int {
	var res int
	total := len(path)
	done := 0
	for _, i := range path {
		done++
		y := i / grid.width
		x := i % grid.width
		if y == grid.guard.y && x == grid.guard.x {
			continue
		}
		res += simulate(grid.Copy(), y, x)
		fmt.Printf("\r%.2f %% Done.", float64(done)/float64(total)*100)
	}
	return res
}

func simulate(grid Grid, y, x int) int {
	grid.data[y*grid.width+x] = '#'
	visited := make(map[int]int, grid.height*grid.width)
	visited[grid.guard.y*grid.width+grid.guard.x] |= int(grid.guard.facing)
	for {
		guradY, guardX := grid.guard.Next()
		offGrid, r := grid.At(guradY, guardX)
		if r == '#' {
			grid.guard.Turn()
			continue
		}

		grid.guard.Step()
		if visited[grid.guard.y*grid.width+grid.guard.x]&int(grid.guard.facing) != 0 {
			return 1
		}
		visited[grid.guard.y*grid.width+grid.guard.x] |= int(grid.guard.facing)
		if !offGrid {
			return 0
		}
	}
}

func part1(grid Grid) {
	for {
		guradY, guardX := grid.guard.Next()
		offGrid, r := grid.At(guradY, guardX)
		if r == '#' {
			grid.guard.Turn()
			continue
		}
		grid.data[grid.guard.y*grid.width+grid.guard.x] = 'X'
		grid.guard.Step()
		if !offGrid {
			return
		}
	}
}

func parse(content string) Grid {
	lines := strings.Split(content, "\n")
	height, width := len(lines)-1, len(lines[0])
	data := make([]rune, height*width)
	var guardX, guardY int
	for y, line := range lines {
		for x, char := range line {
			data[y*width+x] = char
			if char == '^' {
				guardY = y
				guardX = x
				data[y*width+x] = '.'
			}
		}
	}
	guard := Guard{y: guardY, x: guardX, facing: Up}
	return Grid{data: data, height: height, width: width, guard: guard}

}

func (g Grid) At(y, x int) (bool, rune) {
	if y < 0 || y >= g.height || x < 0 || x >= g.width {
		return false, rune(0)
	}
	return true, g.data[y*g.width+x]
}

func (g Grid) Count() int {
	return len(g.Path())
}

func (g Grid) Path() []int {
	var path []int
	for y := range g.height {
		for x := range g.width {
			_, r := g.At(y, x)
			if r == 'X' {
				path = append(path, y*g.width+x)
			}
		}
	}
	return path
}

func (g Grid) String() string {
	b := strings.Builder{}
	for y := range g.height {
		for x := range g.width {
			_, char := g.At(y, x)
			if y == g.guard.y && x == g.guard.x {
				char = g.guard.Rune()
			}
			fmt.Fprintf(&b, string(char))
		}
		fmt.Fprintln(&b)
	}
	return b.String()
}

func (g Guard) Next() (int, int) {
	switch g.facing {
	case Up:
		g.y--
	case Right:
		g.x++
	case Down:
		g.y++
	case Left:
		g.x--
	}
	return g.y, g.x
}

func (g *Guard) Step() {
	switch g.facing {
	case Up:
		g.y--
	case Right:
		g.x++
	case Down:
		g.y++
	case Left:
		g.x--
	}
}

func (g *Guard) Turn() {
	switch g.facing {
	case Up:
		g.facing = Right
	case Right:
		g.facing = Down
	case Down:
		g.facing = Left
	case Left:
		g.facing = Up
	}
}

func (g *Guard) Rune() rune {
	switch g.facing {
	case Up:
		return '^'
	case Right:
		return '>'
	case Down:
		return 'v'
	case Left:
		return '<'
	default:
		panic(fmt.Sprintf("unknown direction: %d\n", g.facing))
	}
}

func (g Grid) Copy() Grid {
	newData := make([]rune, len(g.data))
	copy(newData, g.data)
	newGuard := g.guard
	return Grid{
		data:   newData,
		height: g.height,
		width:  g.width,
		guard:  newGuard,
	}
}
