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

type Circuit map[int]struct{}

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

	subset := append([]Distance(nil), distances[:1000]...)
	circuits := createCircuits(subset)
	slices.SortFunc(circuits, func(a, b Circuit) int {
		return cmp.Compare(len(b), len(a))
	})
	res1 := part1(circuits[:3])
	fmt.Printf("Part1: %d\n", res1)

	lastDist := largeCircuit(distances, len(boxes))
	res2 := part2(boxes[lastDist.fromIdx], boxes[lastDist.toIdx])
	fmt.Printf("Part2: %d\n", res2)
}

func part2(box1, box2 Box) int {
	return box1.x * box2.x
}

func largeCircuit(distances []Distance, boxesCount int) Distance {
	looseCircuits := []Circuit{}
	lastDist := Distance{}
	for {
		if len(looseCircuits) == 1 && len(looseCircuits[0]) == boxesCount {
			break
		}
		lastDist = distances[0]
		distances = distances[1:]

		circuit := Circuit{lastDist.fromIdx: {}, lastDist.toIdx: {}}

		merged := false
		for _, lc := range looseCircuits {
			for idx := range circuit {
				if _, ok := lc[idx]; ok {
					merged = true
					break
				}
			}
			if merged {
				for idx := range circuit {
					lc[idx] = struct{}{}
				}
				break
			}
		}
		if merged {
		outer:
			for {
				for i, lc1 := range looseCircuits {
					for j, lc2 := range looseCircuits {
						if i == j || j < i {
							continue
						}
						needMerge := false
						for idx := range lc1 {
							if _, ok := lc2[idx]; ok {
								needMerge = true
								break
							}
						}
						if needMerge {
							for idx := range lc2 {
								lc1[idx] = struct{}{}
							}
							// remove j
							looseCircuits = append(looseCircuits[:j], looseCircuits[j+1:]...)
							continue outer
						}
					}
				}
				break
			}
		} else {
			looseCircuits = append(looseCircuits, circuit)
		}
	}

	return lastDist
}

func part1(curcuits []Circuit) int {
	res := 1
	for _, c := range curcuits {
		res *= len(c)
	}
	return res
}

func createCircuits(distances []Distance) []Circuit {
	circuits := []Circuit{}
	for len(distances) > 0 {
		dist := distances[0]
		distances = distances[1:]

		circuit := Circuit{dist.fromIdx: {}, dist.toIdx: {}}
	outer:
		for {
			for i, d := range distances {
				_, fromOk := circuit[d.fromIdx]
				_, toOk := circuit[d.toIdx]
				if fromOk || toOk {
					circuit[d.fromIdx] = struct{}{}
					circuit[d.toIdx] = struct{}{}
					distances = append(distances[:i], distances[i+1:]...)
					continue outer // restart searching
				}
			}
			// no more connections found
			break
		}
		circuits = append(circuits, circuit)
	}

	return circuits
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
