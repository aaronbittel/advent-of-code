package main

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	y int
	x int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <input-file>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	points := parse(f)

	allAreas := allAreas(points)
	slices.SortFunc(allAreas, func(a, b Area) int {
		return cmp.Compare(b.area, a.area)
	})

	res1 := allAreas[0].area
	fmt.Printf("Part1: %d\n", res1)

	polygon := Polygon(points)
	res2 := part2(allAreas, polygon)
	fmt.Printf("Part2: %d\n", res2)
}

type Area struct {
	p1   Point
	p2   Point
	area int
}

func allAreas(points []Point) []Area {
	areas := []Area{}
	for i, p1 := range points {
		for j, p2 := range points {
			if i == j || i < j {
				continue
			}
			areas = append(areas, Area{p1: p1, p2: p2, area: area(p1, p2)})
		}
	}
	return areas
}

// working, but slow solution
func part2(areas []Area, polygon Polygon) int {
	var maxArea int
outer:
	for i, a := range areas {
		fmt.Printf("\rChecking (%d/%d)", i+1, len(areas))
		p1 := a.p1
		p2 := a.p2
		lu := p1
		rl := p2
		if p1.x < p2.x {
			if p1.y > p2.y {
				lu = Point{y: min(p1.y, p2.y), x: min(p1.x, p2.x)}
				rl = Point{y: max(p1.y, p2.y), x: max(p1.x, p2.x)}
			}
		} else {
			if p1.y < p2.y {
				rl = Point{y: max(p1.y, p2.y), x: max(p1.x, p2.x)}
				lu = Point{y: min(p1.y, p2.y), x: min(p1.x, p2.x)}
			} else {
				lu, rl = rl, lu
			}

		}
		for x := lu.x; x <= rl.x; x++ {
			if !polygon.Contains(lu.y, x) || !polygon.Contains(rl.y, x) {
				continue outer
			}
		}
		for y := lu.y + 1; y < rl.y; y++ {
			if !polygon.Contains(y, lu.x) || !polygon.Contains(y, rl.x) {
				continue outer
			}
		}

		maxArea = area(p1, p2)
		break
	}
	fmt.Printf("\r\033[2K")
	return maxArea
}

type Tiles map[Point]struct{}

func (t Tiles) Contains(y, x int) bool {
	_, ok := t[Point{y: y, x: x}]
	if ok {
		return true
	}
	var up, down, left, right bool

	for p := range t {
		if p.x == x {
			if p.y < y {
				up = true
			} else {
				down = true
			}
			if up && down {
				return true
			}
		}
		if p.y == y {
			if p.x < x {
				left = true
			} else {
				right = true
			}
			if left && right {
				return true
			}
		}
	}
	return false
}

func (t Tiles) AddPoints(p1, p2 Point) {
	if p1.y == p2.y {
		minX, maxX := min(p1.x, p2.x), max(p1.x, p2.x)
		for x := minX; x <= maxX; x++ {
			t[Point{y: p1.y, x: x}] = struct{}{}
		}
		return
	}
	minY, maxY := min(p1.y, p2.y), max(p1.y, p2.y)
	for y := minY; y <= maxY; y++ {
		t[Point{y: y, x: p1.x}] = struct{}{}
	}
}

func parseTiles(points []Point) Tiles {
	tiles := Tiles{}
	p1 := points[len(points)-1]
	for i := 0; i < len(points); i++ {
		p2 := points[i]
		tiles.AddPoints(p1, p2)
		p1 = p2
	}
	return tiles
}

func area(p1, p2 Point) int {
	dy := math.Abs(float64(p1.y-p2.y)) + 1
	dx := math.Abs(float64(p1.x-p2.x)) + 1
	return int(dy * dx)
}

func parse(f io.Reader) []Point {
	points := []Point{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		points = append(points, Point{y: y, x: x})
	}
	return points
}

func (t Tiles) String() string {
	var (
		minX = math.MaxInt
		maxX = 0
		minY = math.MaxInt
		maxY = 0
	)

	for p := range t {
		maxY = max(maxY, p.y)
		minY = min(minY, p.y)
		maxX = max(maxX, p.x)
		minX = min(minX, p.x)
	}

	b := strings.Builder{}
	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			if t.Contains(y, x) {
				b.WriteByte('#')
			} else {
				b.WriteByte('.')
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}

type Polygon []Point

func (polygon Polygon) Contains(y, x int) bool {
	numVertices := len(polygon)
	inside := false

	for i := 0; i < numVertices; i++ {
		p1 := polygon[i]
		p2 := polygon[(i+1)%numVertices]

		if (x-p1.x)*(p2.y-p1.y) == (y-p1.y)*(p2.x-p1.x) &&
			x >= min(p1.x, p2.x) && x <= max(p1.x, p2.x) &&
			y >= min(p1.y, p2.y) && y <= max(p1.y, p2.y) {
			return true
		}

		if (p1.y > y) != (p2.y > y) {
			xIntersection := (p2.x-p1.x)*(y-p1.y)/(p2.y-p1.y) + p1.x
			if x < xIntersection {
				inside = !inside
			}
		}
	}

	return inside
}
