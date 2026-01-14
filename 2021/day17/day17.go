package main

import (
	"AOC2021/internal/common"
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"slices"
)

type Rect struct {
	Y      int
	X      int
	Height int
	Width  int
}

type Vector2 struct {
	Y int
	X int
}

func main() {
	f := common.GetFile()
	defer f.Close()

	target := parse(f)

	res, dur := common.TimeIt(func() [2]int {
		res1, res2 := part1(target)
		return [2]int{res1, res2}
	})
	fmt.Printf("Part1: %d, Part2: %d, took %s\n", res[0], res[1], dur)
}

func part1(target Rect) (int, int) {
	var (
		count int
		res   int
		minY  = target.Y - target.Height
		maxY  = intAbs(minY)
		maxX  = target.X + target.Width
	)
	for y := minY; y < maxY; y++ {
		for x := 1; x < maxX; x++ {
			height, ok := target.Try(y, x)
			if ok {
				res = max(res, height)
				count++
			}
		}
	}
	return res, count
}

func (r Rect) Try(y, x int) (int, bool) {
	var (
		pos    = Vector2{}
		vel    = Vector2{Y: y, X: x}
		height int
	)

	for !r.Beyond(pos.Y, pos.X) {
		if r.Inside(pos.Y, pos.X) {
			return height, true
		}
		pos, vel = pos.Advance(vel.Y, vel.X)
		height = max(height, pos.Y)
	}

	return 0, false
}

func parse(r io.Reader) Rect {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		panic("illegal input")
	}
	var (
		x1, x2, y1, y2 int
	)
	n, err := fmt.Sscanf(scanner.Text(), "target area: x=%d..%d, y=%d..%d", &x1, &x2, &y1, &y2)
	if err != nil {
		log.Fatal(err)
	}
	if n != 4 {
		log.Fatalf("expected 4 values, but got %d", n)
	}
	return Rect{
		Y:      max(y1, y2),
		X:      min(x1, x2),
		Height: intAbs(y1-y2) + 1,
		Width:  intAbs(x1-x2) + 1,
	}
}

func (r Rect) Inside(y, x int) bool {
	return y <= r.Y && y >= r.Y-r.Height+1 && x >= r.X && x <= r.X+r.Width-1
}

func (r Rect) Beyond(y, x int) bool {
	return r.TooShort(y, x) || r.TooFar(y, x)
}

func (r Rect) TooShort(y, x int) bool {
	return y < r.Y-r.Height-1
}

func (r Rect) TooFar(y, x int) bool {
	return x > r.X+r.Width-1
}

func (v Vector2) Advance(y, x int) (Vector2, Vector2) {
	newPos := Vector2{
		Y: v.Y + y,
		X: v.X + x,
	}
	var newVelX int
	if x > 0 {
		newVelX = x - 1
	} else if x < 0 {
		newVelX = x + 1
	}

	newVel := Vector2{
		Y: y - 1,
		X: newVelX,
	}

	return newPos, newVel
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func debug(target Rect, velocity Vector2) {
	var (
		pos       = Vector2{}
		vel       = velocity
		positions = []Vector2{pos}
	)

	for !target.Beyond(pos.Y, pos.X) {
		if target.Inside(pos.Y, pos.X) {
			break
		}
		pos, vel = pos.Advance(vel.Y, vel.X)
		positions = append(positions, pos)
	}

	minY, maxY, minX, maxX := bounds(positions)
	minY = min(minY, target.Y-target.Height)
	maxX = max(maxX, target.X+target.Width)

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			if y == 0 && x == 0 {
				fmt.Print("S")
				continue
			}
			if slices.Contains(positions, Vector2{Y: y, X: x}) {
				fmt.Print("#")
			} else if target.Inside(y, x) {
				fmt.Print("T")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func bounds(positions []Vector2) (int, int, int, int) {
	var (
		minY int = math.MaxInt
		maxY int = math.MinInt
		minX int = math.MaxInt
		maxX int = math.MinInt
	)

	for _, pos := range positions {
		minY = min(minY, pos.Y)
		maxY = max(maxY, pos.Y)
		minX = min(minX, pos.X)
		maxX = max(maxX, pos.X)
	}

	return minY, maxY, minX, maxX
}
