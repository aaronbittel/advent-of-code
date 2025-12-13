package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Device map[string][]string

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	device := parse(f)

	res1 := part1(device)
	fmt.Printf("Part1: %d\n", res1)

	res2 := part2(device)
	fmt.Printf("Part2: %d\n", res2)

}

func traverse(device Device, start, end string, visited map[string]struct{}, scores map[string]int) int {
	if start == end {
		return 1
	}
	_, ok := visited[start]
	if ok || start == "out" {
		return 0
	}
	if s, ok := scores[start]; ok {
		return s
	}

	visited[start] = struct{}{}
	var total int
	for _, node := range device[start] {
		total += traverse(device, node, end, visited, scores)
	}
	delete(visited, start)
	scores[start] = total
	return total
}

func part2(device Device) int {
	a1 := traverse(device, "svr", "fft", make(map[string]struct{}), make(map[string]int))
	a2 := traverse(device, "fft", "dac", make(map[string]struct{}), make(map[string]int))
	a3 := traverse(device, "dac", "out", make(map[string]struct{}), make(map[string]int))
	b1 := traverse(device, "svr", "dac", make(map[string]struct{}), make(map[string]int))
	b2 := traverse(device, "dac", "fft", make(map[string]struct{}), make(map[string]int))
	b3 := traverse(device, "fft", "out", make(map[string]struct{}), make(map[string]int))
	return a1*a2*a3 + b1*b2*b3
}

func part1(device Device) int {
	visited := make(map[string]struct{})
	scores := make(map[string]int)
	return traverse(device, "you", "out", visited, scores)
}

func parse(f io.Reader) Device {
	device := Device{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if _, ok := device[parts[0]]; ok {
			panic("already in map")
		}
		device[parts[0]] = strings.Split(parts[1], " ")
	}
	return device
}
