package main

import (
	"AOC2021/internal/common"
	"bufio"
	"fmt"
	"io"
	"maps"
	"strings"
)

const (
	Start = "start"
	End   = "end"
)

type Graph map[string][]string

func main() {
	f := common.GetFile()
	defer f.Close()

	graph := parse(f)

	res1, dur1 := common.TimeIt(func() int {
		return part1(graph)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)
}

func part1(graph Graph) int {
	visited := map[string]struct{}{Start: {}}
	return traverse(graph, Start, visited)
}

func traverse(graph Graph, pos string, visitedSmallCaves map[string]struct{}) int {
	var count int
	for _, neighbour := range graph[pos] {
		clonedMap := maps.Clone(visitedSmallCaves)
		if _, ok := visitedSmallCaves[neighbour]; ok {
			continue
		}
		if isSmallCave(neighbour) {
			clonedMap[neighbour] = struct{}{}
		}
		if neighbour == End {
			count++
			continue
		}
		count += traverse(graph, neighbour, clonedMap)
	}
	return count
}

func isSmallCave(node string) bool {
	return node[0] >= 97 && node[0] <= 122
}

func parse(r io.Reader) Graph {
	graph := Graph{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")
		from, to := parts[0], parts[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	return graph
}
