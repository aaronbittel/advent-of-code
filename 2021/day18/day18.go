package main

import (
	"AOC2021/internal/common"
	"bufio"
	"fmt"
	"io"
)

func main() {
	f := common.GetFile()
	defer f.Close()

	snailfishNumbers := parse(f)

	res1, dur1 := common.TimeIt(func() int {
		return part1(snailfishNumbers)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)
}

func part1(snailfishNumbers []*Node) int {
	n := snailfishNumbers[0]
	for _, sn := range snailfishNumbers[1:] {
		n = n.Sum(sn)
	}
	return n.Magnitude()
}

func parse(r io.Reader) []*Node {
	nums := []*Node{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		nums = append(nums, ParseNode(scanner.Text()))
	}
	return nums
}
