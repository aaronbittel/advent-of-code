package main

import (
	"AOC2022/internal/common"
	"bufio"
	"fmt"
	"io"
	"log"
	"maps"
	"math"
	"strings"
)

type Point struct {
	Y, X int
}

type Cell byte

const (
	ROCK  Cell = '#'
	AIR   Cell = '.'
	SAND  Cell = 'o'
	START Cell = '+'
)

type Grid map[Point]Cell

var START_POINT = Point{Y: 0, X: 500}

func main() {
	f := common.GetFile()
	defer f.Close()

	grid := parse(f)

	part1Grid := maps.Clone(grid)
	res1, dur1 := common.TimeIt(func() int {
		return part1(part1Grid)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)

	res2, dur2 := common.TimeIt(func() int {
		return part2(grid)
	})
	fmt.Printf("Part2: %d, took %s\n", res2, dur2)
}

func part2(grid Grid) int {
	var groundLevel = math.MinInt
	for point := range grid {
		groundLevel = max(groundLevel, point.Y)
	}

	var (
		res   int
		point Point
	)
	for res = 0; point != START_POINT; res++ {
		point, _ = grid.DropOne(groundLevel+1, false)
	}
	return res
}

func part1(grid Grid) int {
	var groundLevel = math.MinInt
	for point := range grid {
		groundLevel = max(groundLevel, point.Y)
	}

	var res int
	for {
		_, done := grid.DropOne(groundLevel, true)
		if done {
			break
		}
		res++
	}
	return res
}

func (g Grid) DropOne(groundLevel int, part1 bool) (Point, bool) {
	var (
		point   = START_POINT
		settled bool
	)
	for {
		point, settled = g.NextPoint(point)
		if settled {
			g[point] = SAND
			return point, false
		}
		if part1 && point.Y > groundLevel {
			return point, true
		}
		if !part1 && point.Y == groundLevel {
			g[point] = SAND
			return point, false
		}
	}
}

var FALLING_DIRECTIONS = []Point{{1, 0}, {1, -1}, {1, 1}}

func (g Grid) NextPoint(point Point) (Point, bool) {
	for _, dir := range FALLING_DIRECTIONS {
		newPoint := Point{Y: point.Y + dir.Y, X: point.X + dir.X}
		if _, ok := g[newPoint]; !ok {
			return newPoint, false
		}
	}

	return point, true
}

func parse(r io.Reader) Grid {
	grid := Grid{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var last Point
		for i, pointStr := range strings.Split(scanner.Text(), " -> ") {
			var y, x int
			n, err := fmt.Sscanf(pointStr, "%d,%d", &x, &y)
			if err != nil {
				log.Fatal(err)
			}
			if n != 2 {
				log.Fatal("invalid input")
			}
			point := Point{Y: y, X: x}
			if i == 0 {
				last = point
				continue
			}
			grid.AddPoints(last, point)
			last = point
		}
	}

	grid[START_POINT] = START

	return grid
}

func (g Grid) AddPoints(last, point Point) {
	if last.Y == point.Y {
		if last.X > point.X {
			last, point = point, last
		}
		for x := last.X; x <= point.X; x++ {
			g[Point{Y: last.Y, X: x}] = ROCK
		}
	} else {
		if last.Y > point.Y {
			last, point = point, last
		}
		for y := last.Y; y <= point.Y; y++ {
			g[Point{Y: y, X: last.X}] = ROCK
		}
	}
}

func (g Grid) String() string {
	var (
		minY = math.MaxInt
		minX = math.MaxInt
		maxY = math.MinInt
		maxX = math.MinInt

		offsetX = 5
		offsetY = 2
	)
	for p := range g {
		minY = min(minY, p.Y)
		minX = min(minX, p.X)
		maxY = max(maxY, p.Y)
		maxX = max(maxX, p.X)
	}

	sb := strings.Builder{}
	sb.Grow((maxY + offsetY - minY) * (maxX + offsetX - minX))
	for y := minY - offsetY; y <= maxY+offsetY; y++ {
		for x := minX - offsetX; x <= maxX+offsetX; x++ {
			if ch, ok := g[Point{Y: y, X: x}]; ok {
				sb.WriteByte(byte(ch))
			} else {
				sb.WriteByte(byte(AIR))
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}
