package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	content := string(data)
	part1DiskMap := parsePart1(content)

	res1 := part1(part1DiskMap)
	fmt.Printf("Part1: %d\n", res1)

	diskMap := parsePart2(content)
	res2 := moveFiles(diskMap)
	fmt.Printf("Part2: %d\n", res2)

}

type Block struct {
	pos  int
	size int
}

func moveFiles(diskMap []int) int {
	files := make(map[int]Block)
	var spaces []Block
	pos := 0
	for i, n := range diskMap {
		if i%2 == 1 {
			spaces = append(spaces, Block{pos: pos, size: n})
		} else {
			files[i/2] = Block{pos: pos, size: n}
		}
		pos += n
	}

	sortedIds := make([]int, len(files), len(files))
	i := 0
	for id := range maps.Keys(files) {
		sortedIds[i] = id
		i++
	}

	sort.Slice(sortedIds, func(i, j int) bool {
		return sortedIds[i] > sortedIds[j]
	})

	for _, id := range sortedIds {
		file := files[id]
		for j, space := range spaces {
			if space.pos >= file.pos {
				break
			}
			if file.size > space.size {
				continue
			}
			files[id] = Block{pos: space.pos, size: file.size}
			newSpaceSize := space.size - file.size
			if newSpaceSize == 0 {
				spaces = append(spaces[:j], spaces[j+1:]...)
			} else {
				spaces[j] = Block{pos: space.pos + file.size, size: newSpaceSize}
			}
			break
		}

	}

	var res int
	for id, file := range files {
		for n := file.pos; n < file.pos+file.size; n++ {
			res += id * n
		}
	}

	return res
}

func part1(diskMap []int) int {
	var (
		freeIdx = nextFreeIdx(diskMap, 0)
		fileIdx = nextFileIdx(diskMap, len(diskMap)-1)
	)

	for diskMap[fileIdx] == -1 {
		fileIdx--
	}

	for freeIdx < fileIdx {
		diskMap[freeIdx], diskMap[fileIdx] = diskMap[fileIdx], diskMap[freeIdx]
		freeIdx = nextFreeIdx(diskMap, freeIdx)
		fileIdx = nextFileIdx(diskMap, fileIdx)
	}

	var res int
	for i, n := range diskMap {
		if n == -1 {
			break
		}
		res += i * n
	}
	return res
}

func parsePart2(content string) []int {
	var diskMap []int

	for _, char := range content[:len(content)-1] {
		diskMap = append(diskMap, int(char)-'0')
	}

	return diskMap
}

func parsePart1(content string) []int {
	var diskMap []int

	for i, char := range content {
		num := int(char) - '0'
		v := i / 2
		if i%2 == 1 {
			v = -1
		}
		for range num {
			diskMap = append(diskMap, v)
		}
	}

	return diskMap
}

func nextFreeIdx(diskMap []int, idx int) int {
	for diskMap[idx] != -1 {
		idx++
	}
	return idx
}

func nextFileIdx(diskMap []int, idx int) int {
	for diskMap[idx] == -1 {
		idx--
	}
	return idx
}
