package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Grid struct {
	data   []int
	height int
	width  int
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
	grid, startingIdx := parse(content)

	startPart1 := time.Now()
	res1 := part1(grid, startingIdx)
	part1Dur := time.Since(startPart1)
	fmt.Printf("Part1: %d, took: %s\n", res1, part1Dur)

	startPart2 := time.Now()
	res2 := part2(grid, startingIdx)
	part2Dur := time.Since(startPart2)
	fmt.Printf("Part2: %d, took: %s\n", res2, part2Dur)
}

func part2(grid Grid, startingIdx []int) int {
	var res int

	for _, idx := range startingIdx {
		paths := traverse(grid, idx)
		res += len(paths)
	}

	return res
}

func part1(grid Grid, startingIdx []int) int {
	uniquePoints := func(points []Point) int {
		unique := make(map[Point]struct{})
		for _, point := range points {
			unique[point] = struct{}{}
		}
		return len(unique)
	}

	var res int

	for _, idx := range startingIdx {
		res += uniquePoints(traverse(grid, idx))
	}

	return res
}

type Point struct {
	y int
	x int
}

func traverse(grid Grid, start int) []Point {
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var found []Point
	y, x := grid.intToCoord(start)
	queue := []Point{Point{y, x}}

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]

		_, curNum := grid.At(point.y, point.x)

		if curNum == 9 {
			found = append(found, point)
			continue
		}

		for _, dir := range directions {
			dy, dx := dir[0], dir[1]
			newY, newX := point.y+dy, point.x+dx
			ok, num := grid.At(newY, newX)
			if !ok {
				continue
			}
			if num-curNum != 1 {
				continue
			}
			queue = append(queue, Point{newY, newX})
		}

	}

	return found
}

func parse(content string) (Grid, []int) {
	lines := strings.Split(content, "\n")
	height, width := len(lines)-1, len(lines[0])

	data := make([]int, height*width)
	var startingIdx []int
	for y, line := range lines {
		for x, char := range line {
			idx := y*width + x
			value := int(char) - 48
			data[idx] = value
			if value == 0 {
				startingIdx = append(startingIdx, idx)
			}
		}
	}

	return Grid{
		data:   data,
		height: height,
		width:  width,
	}, startingIdx
}

func (g Grid) At(y, x int) (bool, int) {
	if y < 0 || y >= g.height || x < 0 || x >= g.width {
		return false, 0
	}
	return true, g.data[y*g.width+x]
}

func (g Grid) String() string {
	b := strings.Builder{}

	for y := range g.height {
		for x := range g.width {
			idx := y*g.width + x
			fmt.Fprintf(&b, strconv.Itoa(g.data[idx]))
		}
		fmt.Fprintln(&b)
	}

	return b.String()
}

func (g Grid) intToCoord(x int) (int, int) {
	return x / g.width, x % g.width
}
