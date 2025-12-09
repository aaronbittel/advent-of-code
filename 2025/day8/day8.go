package main

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Box struct {
	x int
	y int
	z int
}

type Distance struct {
	dist    float64
	fromIdx int
	toIdx   int
}

type DSU struct {
	parents []int
	sizes   []int
}

func NewDSU(size int) DSU {
	parents := make([]int, size)
	sizes := make([]int, size)
	for i := range size {
		parents[i] = i
		sizes[i] = 1
	}
	return DSU{
		parents: parents,
		sizes:   sizes,
	}
}

func (dsu *DSU) find(x int) int {
	parent := dsu.parents[x]
	if parent == x {
		return x
	}
	return dsu.find(parent)
}

func (dsu *DSU) union(x, y int) int {
	parent1 := dsu.find(x)
	parent2 := dsu.find(y)
	if parent1 == parent2 {
		return dsu.sizes[parent1]
	}
	if dsu.sizes[parent1] > dsu.sizes[parent2] {
		dsu.parents[parent2] = parent1
		dsu.sizes[parent1] += dsu.sizes[parent2]
		return dsu.sizes[parent1]
	}
	dsu.parents[parent1] = parent2
	dsu.sizes[parent2] += dsu.sizes[parent1]
	return dsu.sizes[parent2]

}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <input-file>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	boxes := parse(f)
	distances := createDistanceGrid(boxes)
	slices.SortFunc(distances, func(a, b Distance) int {
		return cmp.Compare(a.dist, b.dist)
	})

	res1 := part1(distances, len(boxes), 1000)
	fmt.Printf("Part1: %d\n", res1)

	res2 := part2(boxes, distances)
	fmt.Printf("Part2: %d\n", res2)
}

func part2(boxes []Box, distances []Distance) int {
	lastDist := Distance{}
	dsu := NewDSU(len(boxes))
	for _, d := range distances {
		if dsu.union(d.fromIdx, d.toIdx) == len(boxes) {
			lastDist = d
			break
		}
	}
	return boxes[lastDist.fromIdx].x * boxes[lastDist.toIdx].x
}

func part1(distances []Distance, size, count int) int {
	dsu := NewDSU(size)
	for _, d := range distances[:count] {
		dsu.union(d.fromIdx, d.toIdx)
	}

	sizes := append([]int{}, dsu.sizes...)
	slices.SortFunc(sizes, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	res := 1
	for _, s := range sizes[:3] {
		res *= s
	}
	return res
}

func createDistanceGrid(boxes []Box) []Distance {
	distances := make([]Distance, 0, len(boxes)*len(boxes))
	for i, ibox := range boxes {
		for j, jbox := range boxes {
			if i == j || j < i {
				continue
			}
			distances = append(distances, Distance{
				dist:    ibox.Distance(jbox),
				fromIdx: i,
				toIdx:   j,
			})
		}
	}
	return distances
}

func parse(f io.Reader) []Box {
	boxes := []Box{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")
		box := BoxFromSlice(coords)
		boxes = append(boxes, box)
	}
	return boxes
}

func (b Box) Distance(other Box) float64 {
	dx := float64(b.x - other.x)
	dy := float64(b.y - other.y)
	dz := float64(b.z - other.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func BoxFromSlice(coords []string) Box {
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		log.Fatal(err)
	}
	z, err := strconv.Atoi(coords[2])
	if err != nil {
		log.Fatal(err)
	}
	return Box{
		x: x,
		y: y,
		z: z,
	}
}
