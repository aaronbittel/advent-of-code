package main

import (
	"AOC2021/internal/common"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type LavaMap [][]int

type Point struct {
	Y int
	X int
}

func main() {
	f := common.GetFilename()

	lavaMap := parse(f)
	lowPoints := getLowPoints(lavaMap)

	part1, dur1 := common.TimeIt(func() int {
		return part1(lavaMap, lowPoints)
	})
	fmt.Printf("Part1: %d, took %s\n", part1, dur1)

	part2, dur2 := common.TimeIt(func() int {
		return part2(lavaMap)
	})
	fmt.Printf("Part2: %d, took %s\n", part2, dur2)
}

func part2(lavaMap LavaMap) int {
	points := make(map[Point]int, lavaMap.Height()*lavaMap.Width())
	for y, row := range lavaMap {
		for x, cell := range row {
			if cell != 9 {
				points[Point{Y: y, X: x}] = cell
			}
		}
	}

	results := []int{}
	for len(points) > 0 {
		results = append(results, basin(points))
	}
	slices.Sort(results)
	last := len(results) - 1
	return results[last] * results[last-1] * results[last-2]
}

func anyPoint(points map[Point]int) Point {
	for k := range points {
		return k
	}
	panic("map should not be empty")
}

func basin(points map[Point]int) int {
	var res int

	start := anyPoint(points)

	queue := make([]Point, 0, 64)
	queue = append(queue, start)
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]

		_, ok := points[point]
		if !ok {
			continue
		}
		res++
		delete(points, point)

		for _, dir := range DIRECTIONS {
			newPoint := Point{Y: point.Y + dir.Y, X: point.X + dir.X}
			if _, nOk := points[newPoint]; nOk {
				queue = append(queue, newPoint)
			}
		}
	}

	return res
}

func part1(lavaMap LavaMap, lowPoints []Point) int {
	var res int
	for _, lp := range lowPoints {
		res += lavaMap[lp.Y][lp.X] + 1
	}
	return res
}

func getLowPoints(lavaMap LavaMap) []Point {
	lowPoints := []Point{}
	for y, row := range lavaMap {
		for x := range row {
			if lavaMap.LowPoint(y, x) {
				lowPoints = append(lowPoints, Point{y, x})
			}
		}
	}
	return lowPoints
}

var DIRECTIONS = []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func (l LavaMap) LowPoint(y, x int) bool {
	v := l[y][x]
	for _, dir := range DIRECTIONS {
		newY, newX := y+dir.Y, x+dir.X
		if !l.InBounds(newY, newX) {
			continue
		}
		if l[newY][newX] <= v {
			return false
		}
	}
	return true
}

func (l LavaMap) Width() int {
	return len(l[0])
}

func (l LavaMap) Height() int {
	return len(l)
}

func parse(filename string) LavaMap {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	content := string(data)
	lines := strings.Split(content, "\n")
	lines = lines[:len(lines)-1] // skip last empty line

	lavaMap := make([][]int, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "")
		il := make([]int, 0, len(parts))
		for _, part := range parts {
			n, err := strconv.Atoi(part)
			if err != nil {
				log.Fatal(err)
			}
			il = append(il, n)
		}
		lavaMap = append(lavaMap, il)
	}
	return lavaMap
}

func (l LavaMap) InBounds(y, x int) bool {
	return y >= 0 && y < l.Height() && x >= 0 && x < l.Width()
}

func (l LavaMap) String() string {
	sb := strings.Builder{}
	sb.Grow(l.Height() * (l.Width() + 1))
	for _, row := range l {
		for _, cell := range row {
			sb.WriteString(strconv.Itoa(cell))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (l LavaMap) ShowBasin(basin map[Point]struct{}) string {
	sb := strings.Builder{}
	sb.Grow(l.Height() * (l.Width() + 1))
	for y, row := range l {
		for x, cell := range row {
			if _, ok := basin[Point{y, x}]; ok {
				sb.WriteString(".")
			} else {
				sb.WriteString(strconv.Itoa(cell))
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
