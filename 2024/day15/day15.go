package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"maps"
	"os"
	"strings"
)

type Pos struct {
	y int
	x int
}

type Warehouse struct {
	grid  map[Pos]rune
	robot Pos
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	warehouse, moves := parse(f)
	gridPart1 := maps.Clone(warehouse.grid)
	warehousePart1 := warehouse
	warehousePart1.grid = gridPart1

	res1 := part1(warehousePart1, moves)
	fmt.Printf("Part1: %d\n", res1)

	warehousePart2 := double(warehouse)
	res2 := part2(warehousePart2, moves)
	fmt.Printf("Part2: %d\n", res2)
}

func part2(warehouse Warehouse, moves []rune) int {
	for _, move := range moves {
		switch move {
		case '^':
			warehouse.MoveRobotPart2(-1, 0)
		case '>':
			warehouse.MoveRobotPart2(0, 1)
		case 'v':
			warehouse.MoveRobotPart2(1, 0)
		case '<':
			warehouse.MoveRobotPart2(0, -1)
		default:
			panic("unknown move")
		}
	}

	var res int
	for pos, sym := range warehouse.grid {
		if sym != '[' {
			continue
		}
		res += 100*pos.y + pos.x
	}
	return res
}

func (wh *Warehouse) MoveRobotPart2(y, x int) {
	newPos := Pos{y: wh.robot.y + y, x: wh.robot.x + x}
	oldGrid := maps.Clone(wh.grid)
	switch wh.grid[newPos] {
	case '.':
		wh.grid[newPos] = '@'
		wh.grid[wh.robot] = '.'
		wh.robot = newPos
	case '#':
		return
	case '[':
		if x == 0 {
			if wh.PushBoxesVertically(newPos.y, newPos.x, y, x) {
				wh.MoveRobotPart2(y, x)
			} else {
				wh.grid = oldGrid
			}
		} else {
			if wh.PushBoxesRight(newPos.y, newPos.x+1) {
				wh.MoveRobotPart2(y, x)
			} else {
				wh.grid = oldGrid
			}
		}
	case ']':
		if x == 0 {
			if wh.PushBoxesVertically(newPos.y, newPos.x-1, y, x) {
				wh.MoveRobotPart2(y, x)
			} else {
				wh.grid = oldGrid
			}
		} else {
			if wh.PushBoxesLeft(newPos.y, newPos.x-1) {
				wh.MoveRobotPart2(y, x)
			} else {
				wh.grid = oldGrid
			}
		}
	default:
		panic("invalid symbol")
	}
}

func (wh *Warehouse) PushBoxesLeft(ly, lx int) bool {
	nextPos := Pos{y: ly, x: lx - 1}

	success := false

	if wh.grid[nextPos] == '.' {
		success = true
	} else if wh.grid[nextPos] == '#' {
		return false
	} else if wh.grid[nextPos] == ']' {
		success = wh.PushBoxesLeft(ly, lx-2)
	} else {
		panic(fmt.Sprintf("invalid symbol: %s @ %v", string(wh.grid[nextPos]), nextPos))
	}

	prevPrevPos := Pos{y: ly, x: lx + 1}
	prevPos := Pos{y: ly, x: lx}

	if success {
		wh.grid[prevPrevPos] = '.'
		wh.grid[prevPos] = ']'
		wh.grid[nextPos] = '['
	}

	return success
}
func (wh *Warehouse) PushBoxesRight(ry, rx int) bool {
	nextPos := Pos{y: ry, x: rx + 1}

	success := false

	if wh.grid[nextPos] == '.' {
		success = true
	} else if wh.grid[nextPos] == '#' {
		return false
	} else if wh.grid[nextPos] == '[' {
		success = wh.PushBoxesRight(ry, rx+2)
	} else {
		panic(fmt.Sprintf("invalid symbol: %s", string(wh.grid[nextPos])))
	}

	prevPrevPos := Pos{y: ry, x: rx - 1}
	prevPos := Pos{y: ry, x: rx}

	if success {
		wh.grid[prevPrevPos] = '.'
		wh.grid[prevPos] = '['
		wh.grid[nextPos] = ']'
	}

	return success
}

func (wh *Warehouse) PushBoxesVertically(ly, lx, dy, dx int) bool {
	nextLeft := Pos{y: ly + dy, x: lx + dx}
	nextRight := Pos{y: ly + dy, x: lx + 1 + dx}

	success := false

	if wh.grid[nextLeft] == '.' && wh.grid[nextRight] == '.' {
		success = true
	} else if wh.grid[nextLeft] == '#' || wh.grid[nextRight] == '#' {
		return false
	} else if wh.grid[nextLeft] == '[' {
		success = wh.PushBoxesVertically(nextLeft.y, nextLeft.x, dy, dx)
	} else if wh.grid[nextLeft] == ']' {
		success = wh.PushBoxesVertically(nextLeft.y, nextLeft.x-1, dy, dx)
		if !success {
			return false
		}
		rightPushPos := Pos{y: nextLeft.y, x: nextLeft.x + 1}
		switch wh.grid[rightPushPos] {
		case '.':
			success = true
		case '#':
			return false
		case '[':
			success = success && wh.PushBoxesVertically(rightPushPos.y, rightPushPos.x, dy, dx)
		default:
			panic("impossible push")
		}
	} else if wh.grid[nextRight] == '[' {
		success = wh.PushBoxesVertically(nextRight.y, nextRight.x, dy, dx)
		if !success {
			return false
		}
		leftPushPos := Pos{y: nextLeft.y, x: nextLeft.x}
		switch wh.grid[leftPushPos] {
		case '.':
		case '#':
			return false
		case ']':
			success = success && wh.PushBoxesVertically(leftPushPos.y, leftPushPos.x, dy, dx)
		default:
			panic("impossible push")
		}
	} else {
		panic("impossible move")
	}

	prevLeft := Pos{y: ly, x: lx}
	prevRight := Pos{y: ly, x: lx + 1}

	if success {
		wh.grid[nextLeft] = '['
		wh.grid[prevLeft] = '.'
		wh.grid[nextRight] = ']'
		wh.grid[prevRight] = '.'
	}

	return success
}

func part1(warehouse Warehouse, moves []rune) int {
	for _, move := range moves {
		switch move {
		case '^':
			warehouse.MoveRobotPart1(-1, 0)
		case '>':
			warehouse.MoveRobotPart1(0, 1)
		case 'v':
			warehouse.MoveRobotPart1(1, 0)
		case '<':
			warehouse.MoveRobotPart1(0, -1)
		default:
			panic("unknown move")
		}
	}

	var res int
	for pos, sym := range warehouse.grid {
		if sym != 'O' {
			continue
		}
		res += 100*pos.y + pos.x
	}
	return res
}

func (wh *Warehouse) MoveRobotPart1(y, x int) {
	newPos := Pos{y: wh.robot.y + y, x: wh.robot.x + x}
	switch wh.grid[newPos] {
	case '.':
		wh.grid[newPos] = '@'
		wh.grid[wh.robot] = '.'
		wh.robot = newPos
	case '#':
		return
	case 'O':
		cur := newPos
		for i := 1; ; i++ {
			cur = Pos{y: cur.y + y, x: cur.x + x}
			switch wh.grid[cur] {
			case '.':
				wh.grid[newPos] = '.'
				wh.grid[cur] = 'O'
				wh.MoveRobotPart1(y, x)
				return
			case 'O':
				continue
			case '#':
				return
			default:
				panic(fmt.Sprintf("pushing: invalid symbol '%s'", string(wh.grid[cur])))
			}
		}
	default:
		panic(fmt.Sprintf("looked at: invalid symbol '%s'", string(wh.grid[newPos])))
	}
}

func parse(f io.Reader) (Warehouse, []rune) {
	grid := make(map[Pos]rune)
	var robot Pos

	scanner := bufio.NewScanner(f)
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		if line == "" {
			break
		}
		for x, r := range line {
			if r == '@' {
				robot = Pos{y: y, x: x}
			}
			grid[Pos{y: y, x: x}] = r
		}
	}

	moves := []rune{}
	for scanner.Scan() {
		for _, r := range scanner.Text() {
			moves = append(moves, r)
		}
	}

	return Warehouse{grid: grid, robot: robot}, moves
}

func (wh Warehouse) String() string {
	b := strings.Builder{}

	var maxX, maxY int
	for pos := range wh.grid {
		maxX = max(maxX, pos.x)
		maxY = max(maxY, pos.y)
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			b.WriteRune(wh.grid[Pos{y: y, x: x}])
		}
		b.WriteRune('\n')
	}

	return b.String()
}

func double(warehouse Warehouse) Warehouse {
	newWarehouse := Warehouse{}
	newGrid := make(map[Pos]rune)
	for pos, sym := range warehouse.grid {
		switch sym {
		case '.', '#':
			newGrid[Pos{y: pos.y, x: pos.x * 2}] = sym
			newGrid[Pos{y: pos.y, x: pos.x*2 + 1}] = sym
		case 'O':
			newGrid[Pos{y: pos.y, x: pos.x * 2}] = '['
			newGrid[Pos{y: pos.y, x: pos.x*2 + 1}] = ']'
		case '@':
			robotPos := Pos{y: pos.y, x: pos.x * 2}
			newWarehouse.robot = robotPos
			newGrid[robotPos] = '@'
			newGrid[Pos{y: pos.y, x: pos.x*2 + 1}] = '.'
		}
	}
	newWarehouse.grid = newGrid
	return newWarehouse
}
