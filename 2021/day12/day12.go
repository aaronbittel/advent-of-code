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
		return solve(graph, false)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)

	res2, dur2 := common.TimeIt(func() int {
		return solve(graph, true)
	})
	fmt.Printf("Part2: %d, took %s\n", res2, dur2)
}

func solve(graph Graph, canDoubleVisit bool) int {
	visited := map[string]struct{}{Start: {}}
	return traverse(graph, Start, visited, canDoubleVisit)
}

func traverse(graph Graph, pos string, visitedSmallCaves map[string]struct{}, canDoubleVisit bool) int {
	var count int
	for _, neighbour := range graph[pos] {
		var clonedMap = maps.Clone(visitedSmallCaves)

		// dont move back to start
		if neighbour == Start {
			continue
		}

		// finish at end
		if neighbour == End {
			count++
			continue
		}

		if isSmallCave(neighbour) {
			clonedMap[neighbour] = struct{}{}
			if _, ok := visitedSmallCaves[neighbour]; ok {
				if canDoubleVisit {
					cloneTwiceMap := maps.Clone(visitedSmallCaves)
					count += traverse(graph, neighbour, cloneTwiceMap, false)
				}
				continue
			}
		}

		count += traverse(graph, neighbour, clonedMap, canDoubleVisit)
	}
	return count
}

func isSmallCave(node string) bool {
	return node[0] >= 'a' && node[0] <= 'z'
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
