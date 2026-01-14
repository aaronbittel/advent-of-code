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

	res1, dur1 := common.TimeIt(func() int {
		return part1(target)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)
}

func part1(target Rect) int {
	var (
		res  int
		maxY = intAbs(target.Y - target.Height)
	)
outer:
	for y := 1; y < maxY; y++ {
		for x := 1; ; x++ {
			height, result := target.Try(y, x)
			switch result {
			case Hit:
				res = max(res, height)
			case TooShort:
				continue
			case TooFar:
				continue outer
			}
		}
	}
	return res
}

type Result int

const (
	Hit Result = iota
	TooShort
	TooFar
)

func (r Rect) Try(y, x int) (int, Result) {
	var (
		pos    = Vector2{}
		vel    = Vector2{Y: y, X: x}
		height int
	)

	for {
		if r.Inside(pos.Y, pos.X) {
			return height, Hit
		}
		pos, vel = pos.Advance(vel.Y, vel.X)
		if r.TooShort(pos.Y, pos.X) {
			return 0, TooShort
		}
		if r.TooFar(pos.Y, pos.X) {
			return 0, TooFar
		}
		height = max(height, pos.Y)
	}

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
	yBorder := r.Y - r.Height - 1
	yTrue := y < yBorder
	xBorder := r.X + r.Width - 1
	xTrue := x > xBorder
	return yTrue || xTrue
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
