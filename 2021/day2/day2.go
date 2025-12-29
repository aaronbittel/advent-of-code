package main

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"AOC2021/internal/common"
)

type Command struct {
	Op    string
	Value int
}

func main() {
	f := common.GetFile()
	defer f.Close()

	commands, parseDur := common.TimeIt(func() []Command {
		commands, err := parse(f)
		if err != nil {
			log.Fatal(err)
		}
		return commands
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
		switch cmd.Op {
		case "forward":
			hori += cmd.Value
			depth += aim * cmd.Value
		case "down":
			aim += cmd.Value
		case "up":
			aim -= cmd.Value
		default:
			panic("unknown op")
		}
	}

	return hori * depth
}

func part1(commands []Command) int {
	var hori, depth int

	for _, cmd := range commands {
		switch cmd.Op {
		case "forward":
			hori += cmd.Value
		case "down":
			depth += cmd.Value
		case "up":
			depth -= cmd.Value
		default:
			panic("unknown op")
		}
	}

	return hori * depth
}

func parse(f io.Reader) ([]Command, error) {
	commands := make([]Command, 0, 1024)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var (
			op    string
			value int
		)
		if n, err := fmt.Sscanf(scanner.Text(), "%s %d", &op, &value); n != 2 {
			return nil, err
		}
		commands = append(commands, Command{Op: op, Value: value})
	}

	return commands, scanner.Err()
}
