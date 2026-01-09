package main

import (
	"AOC2021/internal/common"
	"bufio"
	"fmt"
	"io"
	"log"
	"maps"
	"strings"
)

type Point struct {
	Y int
	X int
}

type Instruction struct {
	Horizontal bool
	Count      int
}

type Manual map[Point]struct{}

func main() {
	f := common.GetFile()
	defer f.Close()

	manual, instructions := parse(f)

	manualPart1 := maps.Clone(manual)

	res1, dur1 := common.TimeIt(func() int {
		// only fold once
		return solve(manualPart1, instructions[:1])
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)

	_, dur2 := common.TimeIt(func() int {
		return solve(manual, instructions)
	})
	fmt.Printf("Part2: took %s\n", dur2)
	fmt.Println(manual)
}

func solve(manual Manual, insts []Instruction) int {
	for _, inst := range insts {
		if inst.Horizontal {
			manual.FoldHorizontal(inst.Count)
		} else {
			manual.FoldVertical(inst.Count)
		}
	}
	return len(manual)
}

func (m Manual) FoldHorizontal(y int) {
	for p := range m {
		if p.Y == y {
			panic("point will never be on folding line")
		}
		if p.Y < y {
			continue
		}
		dist := p.Y - y
		newY := y - dist
		// if new points end up in this iteration, they will be skipped, because p.Y < y
		m[Point{Y: newY, X: p.X}] = struct{}{}
		delete(m, p)
	}
}

func (m Manual) FoldVertical(x int) {
	for p := range m {
		if p.X == x {
			panic("point will never be on folding line")
		}
		if p.X < x {
			continue
		}
		dist := p.X - x
		newX := x - dist
		m[Point{Y: p.Y, X: newX}] = struct{}{}
		delete(m, p)
	}
}

func parse(r io.Reader) (Manual, []Instruction) {
	manual := Manual{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var x, y int
		n, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			log.Fatal(err)
		}
		if n != 2 {
			log.Fatal("expected 2 values")
		}
		manual[Point{Y: y, X: x}] = struct{}{}
	}

	instructions := []Instruction{}
	for scanner.Scan() {
		var (
			dir   string
			count int
		)
		n, err := fmt.Sscanf(scanner.Text(), "fold along %1s=%d", &dir, &count)
		if err != nil {
			log.Fatal(err)
		}
		if n != 2 {
			log.Fatal("expected 2 values")
		}
		var horizontal bool
		if dir == "y" {
			horizontal = true
		}
		instructions = append(instructions, Instruction{
			Horizontal: horizontal,
			Count:      count,
		})
	}

	return manual, instructions
}

func (m Manual) String() string {
	sb := strings.Builder{}

	var maxY, maxX int
	for p := range m {
		maxY = max(maxY, p.Y)
		maxX = max(maxX, p.X)
	}

	var p Point
	// inclusive ranges
	for y := range maxY + 1 {
		for x := range maxX + 1 {
			p.Y = y
			p.X = x
			if _, ok := m[p]; ok {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
