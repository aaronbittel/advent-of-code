package main

import (
	"AOC2021/internal/common"
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f := common.GetFile()
	defer f.Close()

	nums := parse(f)

	part1, dur1 := common.TimeIt(func() int {
		return solve(nums, true)
	})
	fmt.Printf("Part1: %d, took %s\n", part1, dur1)

	part2, dur2 := common.TimeIt(func() int {
		return solve(nums, false)
	})
	fmt.Printf("Part2: %d, took %s\n", part2, dur2)
}

func solve(nums []int, constantCost bool) int {
	mi, ma := slices.Min(nums), slices.Max(nums)

	res := math.MaxInt

	for i := mi; i <= ma; i++ {
		res = min(res, calcFuel(nums, i, constantCost))
	}

	return res
}

func calcFuel(nums []int, level int, constantCost bool) int {
	var res int
	for _, num := range nums {
		steps := int(math.Abs(float64(level - num)))
		if constantCost {
			res += steps
		} else {
			res += steps * (steps + 1) / 2
		}
	}
	return res
}

func parse(r io.Reader) []int {
	nums := []int{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		for _, numStr := range strings.Split(scanner.Text(), ",") {
			n, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, n)
		}
	}
	return nums
}
