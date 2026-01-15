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
		snailfishNumbersPart1 := make([]*Node, len(snailfishNumbers))
		for i, sn := range snailfishNumbers {
			snailfishNumbersPart1[i] = sn.Clone()
		}
		return part1(snailfishNumbersPart1)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)

	res2, dur2 := common.TimeIt(func() int {
		return part2(snailfishNumbers)
	})
	fmt.Printf("Part2: %d, took %s\n", res2, dur2)
}

func part2(snailfishNumbers []*Node) int {
	var res int
	for i, n1 := range snailfishNumbers {
		for j, n2 := range snailfishNumbers {
			if i == j {
				continue
			}
			copyN1, copyN2 := n1.Clone(), n2.Clone()
			res = max(res, copyN1.Sum(copyN2).Magnitude())
			copyN1, copyN2 = n1.Clone(), n2.Clone()
			res = max(res, copyN2.Sum(copyN1).Magnitude())
		}
	}
	return res
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
