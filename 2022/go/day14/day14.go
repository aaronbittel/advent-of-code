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

type Grid map[Point]byte

const (
	ROCK  byte = '#'
	AIR   byte = '.'
	SAND  byte = 'o'
	START byte = '+'
)

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
		point = grid.SandFallPart2(groundLevel + 1)
	}
	return res
}

func part1(grid Grid) int {
	var groundLevel = math.MinInt
	for point := range grid {
		groundLevel = max(groundLevel, point.Y)
	}

	var res int
	for res = 0; grid.SandFall(groundLevel); res++ {
	}
	return res
}

func (g Grid) SandFallPart2(groundLevel int) Point {
	var (
		point   = START_POINT
		falling bool
	)
	for {
		point, falling = g.NextPoint(point)
		if !falling || point.Y == groundLevel {
			g[point] = SAND
			return point
		}
	}
}

func (g Grid) SandFall(groundLevel int) bool {
	var (
		point   = START_POINT
		falling bool
	)
	for {
		point, falling = g.NextPoint(point)
		if !falling {
			g[point] = SAND
			return true
		}
		if point.Y > groundLevel {
			return false
		}
	}
}

func (g Grid) NextPoint(p Point) (Point, bool) {
	downPoint := Point{Y: p.Y + 1, X: p.X}
	if _, ok := g[downPoint]; !ok {
		return downPoint, true
	}

	diagLeftPoint := Point{Y: p.Y + 1, X: p.X - 1}
	if _, ok := g[diagLeftPoint]; !ok {
		return diagLeftPoint, true
	}

	diagRightPoint := Point{Y: p.Y + 1, X: p.X + 1}
	if _, ok := g[diagRightPoint]; !ok {
		return diagRightPoint, true
	}

	return p, false
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
				sb.WriteByte(ch)
			} else {
				sb.WriteByte(AIR)
			}
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}
