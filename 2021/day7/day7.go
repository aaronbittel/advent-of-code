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
		return part1(nums)
	})
	fmt.Printf("Part1: %d, took %s\n", part1, dur1)

	part2, dur2 := common.TimeIt(func() int {
		return part2(nums)
	})
	fmt.Printf("Part2: %d, took %s\n", part2, dur2)
}

func part2(nums []int) int {
	mi, ma := slices.Min(nums), slices.Max(nums)

	res := math.MaxInt
	slices.Sort(nums)

	cpy := make([]int, len(nums))

	for i := mi; i <= ma; i++ {
		copy(cpy, nums)
		cost := calcFuel(cpy, i)
		res = min(res, cost)
	}

	return res
}

func calcFuel(nums []int, level int) int {
	var (
		total   int
		curCost = 1
	)

	for len(nums) != 0 {
		var (
			count int
		)
		for i, n := range nums {
			if n == level {
				continue
			}
			count++
			if nums[i] > level {
				nums[i]--
			} else {
				nums[i]++
			}
		}

		total += curCost * count
		curCost++

		newNums := nums[:0]
		for _, n := range nums {
			if n == level {
				continue
			}
			newNums = append(newNums, n)
		}

		nums = newNums
	}

	return total
}

func part1(nums []int) int {
	mi, ma := slices.Min(nums), slices.Max(nums)

	res := math.MaxInt

	for i := mi; i <= ma; i++ {
		res = min(res, calcFuelConstant(nums, i))
	}

	return res
}

func calcFuelConstant(nums []int, level int) int {
	var res int
	for _, num := range nums {
		res += int(math.Abs(float64(level - num)))
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
