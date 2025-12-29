package main

import (
	common "AOC2021/internal"
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type Dir int

const (
	Forward Dir = iota
	Down
	Up
)

type Command struct {
	Dir   Dir
	Value int
}

func main() {
	f := common.GetFile()
	defer f.Close()

	commands, parseDur := common.TimeIt(func() []Command {
		return parse(f)
	})
	fmt.Printf("Parsing took %s\n", parseDur)

	part1, dur1 := common.TimeIt(func() int {
		return part1(commands)
	})
	fmt.Printf("Part1: %d, took %s\n", part1, dur1)

	part2, dur2 := common.TimeIt(func() int {
		return part2(commands)
	})
	fmt.Printf("Part2: %d, took %s\n", part2, dur2)
}

func part2(commands []Command) int {
	var hori, depth, aim int

	for _, cmd := range commands {
		switch cmd.Dir {
		case Forward:
			hori += cmd.Value
			depth += aim * cmd.Value
		case Down:
			aim += cmd.Value
		case Up:
			aim -= cmd.Value
		default:
			panic("unknown dir")
		}
	}

	return hori * depth
}

func part1(commands []Command) int {
	var hori, depth int

	for _, cmd := range commands {
		switch cmd.Dir {
		case Forward:
			hori += cmd.Value
		case Down:
			depth += cmd.Value
		case Up:
			depth -= cmd.Value
		default:
			panic("unknown dir")
		}
	}

	return hori * depth
}

func parse(f io.Reader) []Command {
	commands := []Command{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		dir := NewDir(parts[0])
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		commands = append(commands, Command{Dir: dir, Value: value})
	}

	return commands
}

func NewDir(d string) Dir {
	switch d {
	case "forward":
		return Forward
	case "down":
		return Down
	case "up":
		return Up
	default:
		panic("unknown dir")
	}
}
