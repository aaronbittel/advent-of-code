package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	y int
	x int
}

type Shape map[Point]struct{}

type Region struct {
	width   int
	length  int
	indexes []int
	cur     int
}

type RestorePoint struct {
	shape *Shape
	start Point
}

type RestorePoints []RestorePoint

type Grid struct {
	occupied map[Point]byte
	char     byte

	restorePoints RestorePoints

	width  int
	length int
}

func NewGrid(width, length int) Grid {
	return Grid{
		occupied:      make(map[Point]byte),
		restorePoints: RestorePoints{},
		char:          'A',
		width:         width,
		length:        length,
	}
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

	shapes, regions := parse(string(data))

	res1 := part1(regions, shapes)
	fmt.Printf("Part1: %d\n", res1)
}

func part1(regions []Region, shapes [][]Shape) int {
	resChan := make(chan bool, len(regions))
	sent := 0
	for _, region := range regions {
		if !region.IsDoable(shapes) {
			continue
		}
		sent++
		go func(region Region, shapes [][]Shape, grid Grid) {
			resChan <- FindPlace(region, shapes, grid)
		}(region, shapes, NewGrid(region.width, region.length))
	}

	var res int
	for i := range sent {
		if <-resChan {
			res++
		}
		fmt.Printf("\r(%d / %d)", i+1, sent)
	}
	fmt.Printf("\r\033[2K")
	return res
}

func FindPlace(region Region, shapes [][]Shape, grid Grid) bool {
	idx := region.NextIndex()
	if idx == -1 {
		return true
	}
	rotations := shapes[idx]
	for y := range region.length {
		for x := range region.width {
			if !grid.IsFree(y, x) {
				continue
			}
			for _, rot := range rotations {
				if grid.Place(rot, y, x) {
					if FindPlace(region, shapes, grid) {
						return true
					} else {
						grid.Restore()
					}
				}
			}
		}
	}
	return false
}

func (g *Grid) Place(shape Shape, y, x int) bool {
	// check
	for relP := range shape {
		if !g.IsFree(relP.y+y, relP.x+x) {
			return false
		}
	}

	g.restorePoints = append(g.restorePoints, RestorePoint{
		shape: &shape,
		start: Point{y: y, x: x},
	})

	// commit
	for relP := range shape {
		p := Point{y: relP.y + y, x: relP.x + x}
		g.occupied[p] = g.char
	}

	g.char++
	return true
}

func (g *Grid) Restore() {
	rp := g.restorePoints.Pop()
	if rp == nil {
		return
	}
	for relP := range *rp.shape {
		p := Point{y: relP.y + rp.start.y, x: relP.x + rp.start.x}
		delete(g.occupied, p)
	}
	g.char--
}

func (g Grid) IsFree(y, x int) bool {
	if y < 0 || y >= g.length || x < 0 || x >= g.width {
		return false
	}
	p := Point{y: y, x: x}
	_, ok := g.occupied[p]
	return !ok
}

func (g Grid) At(y, x int) byte {
	if c, ok := g.occupied[Point{y: y, x: x}]; ok {
		return c
	}
	return '.'
}

func (g Grid) String() string {
	b := strings.Builder{}
	for y := range g.length {
		for x := range g.width {
			b.WriteByte(g.At(y, x))
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func (s Shape) Rotations() []Shape {
	rotations := []Shape{s}
	for range 3 {
		last := rotations[len(rotations)-1]
		rotS := Shape{}
		m := Point{y: 1, x: 1}
		if _, ok := last[m]; ok {
			rotS[m] = struct{}{}
		}
		for i := range 3 {
			lp := Point{y: 2 - i, x: 0}
			np := Point{y: 0, x: i}
			if _, ok := last[lp]; ok {
				rotS[np] = struct{}{}
			}
			lp = np
			np = Point{y: i, x: 2}
			if _, ok := last[lp]; ok {
				rotS[np] = struct{}{}
			}
			lp = np
			np = Point{y: 2, x: 2 - i}
			if _, ok := last[lp]; ok {
				rotS[np] = struct{}{}
			}
			lp = np
			np = Point{y: 2 - i, x: 0}
			if _, ok := last[lp]; ok {
				rotS[np] = struct{}{}
			}
		}

		found := false
	outer:
		for _, rot := range rotations {
			for p := range rot {
				if _, ok := rotS[p]; !ok {
					continue outer
				}
			}
			found = true
			break
		}
		if !found {
			rotations = append(rotations, rotS)
		}
	}
	return rotations
}

func (r *Region) NextIndex() int {
	if r.cur >= len(r.indexes) {
		return -1
	}
	idx := r.indexes[r.cur]
	r.cur++
	return idx
}

func parse(content string) ([][]Shape, []Region) {
	parts := strings.Split(content, "\n\n")
	regionParts := strings.Split(parts[len(parts)-1], "\n")

	shapes := make([][]Shape, 0, len(parts)-1)
	regions := make([]Region, 0, len(regionParts))

	// skip region part
	for _, part := range parts[:len(parts)-1] {
		lines := strings.Split(part, "\n")
		// skip index
		shape := Shape{}
		for y, line := range lines[1:] {
			for x, c := range line {
				if c == '#' {
					shape[Point{y: y, x: x}] = struct{}{}
				}
			}
		}
		shapes = append(shapes, shape.Rotations())
	}

	for _, rp := range regionParts[:len(regionParts)-1] {
		p := strings.Split(rp, ": ")
		size := strings.Split(p[0], "x")
		width, err := strconv.Atoi(size[0])
		if err != nil {
			log.Fatal(err)
		}
		length, err := strconv.Atoi(size[1])
		if err != nil {
			log.Fatal(err)
		}

		rawCounts := strings.Split(p[1], " ")
		indexes := make([]int, 0, len(rawCounts))
		for i, rawCount := range rawCounts {
			count, err := strconv.Atoi(rawCount)
			if err != nil {
				log.Fatal(err)
			}
			for range count {
				indexes = append(indexes, i)
			}
		}

		regions = append(regions, Region{
			width: width, length: length, indexes: indexes,
		})
	}

	return shapes, regions
}

func (s Shape) String() string {
	b := strings.Builder{}
	for y := range 3 {
		for x := range 3 {
			p := Point{y: y, x: x}
			if _, ok := s[p]; ok {
				b.WriteByte('#')
			} else {
				b.WriteByte('.')
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func (rps RestorePoints) Pop() *RestorePoint {
	if len(rps) == 0 {
		return nil
	}
	last := len(rps) - 1
	rp := rps[last]
	rps = rps[:last]
	return &rp
}

func (r Region) IsDoable(shapes [][]Shape) bool {
	area := r.width * r.length
	shapeSize := 0
	for {
		idx := r.NextIndex()
		if idx == -1 {
			break
		}
		shapeSize += len(shapes[idx][0])
	}

	r.cur = 0
	return area > shapeSize
}
