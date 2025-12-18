package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type Point struct {
	y int
	x int
}

type Track struct {
	walls map[Point]struct{}

	start Point
	end   Point

	length int
	height int
	width  int
}

type Step struct {
	point  Point
	length int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	track := parse(f)
	path := traverse(track)

	res1, dur1 := timeIt(func() int {
		steps := cheatStepsPart1(track, path)
		return part1(track.length, steps, 100)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)

	res2, dur2 := timeIt(func() int {
		return part2(track, path, 100)
	})
	fmt.Printf("Part2: %d, took %s\n", res2, dur2)
}

func timeIt[T any](f func() T) (T, time.Duration) {
	start := time.Now()
	res := f()
	dur := time.Since(start)
	return res, dur
}

var DIAMOND20 = diamond(20)

func part2(track Track, path map[Point]int, threshold int) int {
	var (
		pos         = track.start
		curDir      Point
		nextDir     Point
		nextPathPos Point
		curLen      int
		res         int
	)

	for pos != track.end {
		var (
			foundNextStep     bool
			visitedCheatPaths bool
		)
		for _, dir := range DIRECTIONS {
			d := Point{y: curDir.y * -1, x: curDir.x * -1}
			if d == dir {
				continue
			}
			newPos := pos.Add(dir)
			if _, ok := track.walls[newPos]; !ok {
				nextDir = dir
				nextPathPos = newPos
				foundNextStep = true
				continue
			}
			if !visitedCheatPaths {
				visitedCheatPaths = true
				for _, step := range DIAMOND20 {
					dir := step.point
					np := pos.Add(dir)
					if v, ok := path[np]; ok && track.length-(curLen+step.length+v) >= threshold {
						res++
					}
				}
			}
			if foundNextStep && visitedCheatPaths {
				break
			}
		}
		curDir = nextDir
		pos = nextPathPos
		curLen++
	}
	return res
}

func diamond(length int) []Step {
	dist := func(p Point) int {
		if p.y < 0 {
			p.y = -p.y
		}
		if p.x < 0 {
			p.x = -p.x
		}
		return p.y + p.x
	}

	positions := make([]Step, 0, length*length)
	left := Point{y: 0, x: -length}
	l := 1
	for range length + 1 {
		for li := range l {
			lp := Point{y: left.y + li, x: left.x}
			positions = append(positions, Step{point: lp, length: dist(lp)})
		}
		left = left.Add(Point{y: -1, x: 1})
		l += 2
	}
	right := Point{y: 0, x: length}
	l = 1
	for range length {
		for li := range l {
			lp := Point{y: right.y + li, x: right.x}
			positions = append(positions, Step{point: lp, length: dist(lp)})
		}
		right = right.Add(Point{y: -1, x: -1})
		l += 2
	}
	return positions
}

func part1(trackLen int, steps []int, saved int) int {
	sort.Ints(steps)
	var res int
	for _, r := range steps {
		if trackLen-r < saved {
			break
		}
		res++
	}
	return res
}

func cheatStepsPart1(track Track, path map[Point]int) []int {
	var (
		pos         = track.start
		curDir      Point
		nextDir     Point
		nextPathPos Point
		curLen      int
		total       = []int{}
	)

	for pos != track.end {
		for _, dir := range DIRECTIONS {
			d := Point{y: curDir.y * -1, x: curDir.x * -1}
			if d == dir {
				continue
			}
			newPos := pos.Add(dir)
			if _, ok := track.walls[newPos]; !ok {
				nextDir = dir
				nextPathPos = newPos
				continue
			}
			// walk through wall
			newPos = newPos.Add(dir)
			if v, ok := path[newPos]; ok {
				// +1 because of the step on path was not counted yet
				// +1 because of the step through the wall
				total = append(total, curLen+v+2)
			}
		}
		curDir = nextDir
		pos = nextPathPos
		curLen++
	}
	return total
}

var DIRECTIONS = []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func traverse(track Track) map[Point]int {
	path := make(map[Point]int)
	pos := track.start
	leftLen := track.length
	var curDir Point

	for pos != track.end {
		for _, dir := range DIRECTIONS {
			d := Point{y: curDir.y * -1, x: curDir.x * -1}
			if d == dir {
				continue
			}
			newPos := pos.Add(dir)
			if _, ok := track.walls[newPos]; ok {
				continue
			}
			leftLen--
			path[newPos] = leftLen
			pos = newPos
			curDir = dir
			break
		}
	}
	return path
}

func parse(f io.Reader) Track {
	track := Track{
		walls: make(map[Point]struct{}),
	}

	scanner := bufio.NewScanner(f)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		track.width = len(line)
		for x, c := range line {
			p := Point{y: y, x: x}
			switch c {
			case '#':
				track.walls[p] = struct{}{}
			case 'E':
				track.length++
				track.end = p
			case 'S':
				track.start = p
			case '.':
				track.length++
			default:
				panic("illegal symbol")
			}
		}
		track.height = y + 1
	}
	return track
}

func (t Track) String() string {
	var maxWidth, maxHeight int
	for p := range t.walls {
		if p.y > maxHeight {
			maxHeight = p.y
		}
		if p.x > maxWidth {
			maxWidth = p.x
		}
	}

	b := strings.Builder{}
	for y := range maxHeight + 1 {
		for x := range maxWidth + 1 {
			p := Point{y: y, x: x}
			if _, ok := t.walls[p]; ok {
				b.WriteByte('#')
			} else if p == t.start {
				b.WriteByte('S')
			} else if p == t.end {
				b.WriteByte('E')
			} else {
				b.WriteByte('.')
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func (p Point) Add(other Point) Point {
	return Point{y: p.y + other.y, x: p.x + other.x}
}

func (t Track) Valid(y, x int) bool {
	if y < 0 || y >= t.height || x < 0 || x >= t.width {
		return false
	}
	return true
}
