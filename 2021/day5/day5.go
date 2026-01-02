package main

import (
	"AOC2021/internal/common"
	"bufio"
	"fmt"
	"io"
	"log"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	Start Point
	End   Point
}

func main() {
	f := common.GetFile()
	defer f.Close()

	lines := parse(f)

	part1, dur1 := common.TimeIt(func() int {
		return solve(lines, false)
	})
	fmt.Printf("Part1: %d, took %s\n", part1, dur1)

	part2, dur2 := common.TimeIt(func() int {
		return solve(lines, true)
	})
	fmt.Printf("Part2: %d, took %s\n", part2, dur2)
}

func solve(lines []Line, useDiagonals bool) int {
	var (
		res  int
		grid = make(map[Point]int)
	)

	for _, line := range lines {
		if !useDiagonals && line.Diagonal() {
			continue
		}

		dir := Point{
			X: sign(line.Start.X - line.End.X),
			Y: sign(line.Start.Y - line.End.Y),
		}

		cur := line.Start
		for {
			grid[cur]++
			if grid[cur] == 2 {
				res++
			}
			if cur == line.End {
				break
			}
			cur.X -= dir.X
			cur.Y -= dir.Y
		}
	}

	return res
}

func parse(f io.Reader) []Line {
	lines := []Line{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var x1, y1, x2, y2 int
		n, err := fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if err != nil {
			log.Fatal(err)
		}
		if n != 4 {
			log.Fatalf("expected 4 nums, but only got %d", n)
		}
		lines = append(lines, NewLine(x1, y1, x2, y2))
	}

	return lines
}

func NewLine(x1, y1, x2, y2 int) Line {
	return Line{
		Start: Point{X: x1, Y: y1},
		End:   Point{X: x2, Y: y2},
	}
}

func (l Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.Start.X, l.Start.Y, l.End.X, l.End.Y)
}

func (l Line) Diagonal() bool {
	return l.Start.X != l.End.X && l.Start.Y != l.End.Y
}

func sign(n int) int {
	return min(1, max(-1, n))
}
