package main

import (
	"AOC2021/internal/common"
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"
)

func main() {
	f := common.GetFile()
	defer f.Close()

	nums := parse(f)

	part1, dur1 := common.TimeIt(func() int {
		part1Nums := make([]int, len(nums))
		copy(part1Nums, nums)
		return solve(part1Nums, 80)
	})
	fmt.Printf("Part1: %d, took %s\n", part1, dur1)

	part2, dur2 := common.TimeIt(func() int {
		return solve(nums, 256)
	})
	fmt.Printf("Part2: %d, took %s\n", part2, dur2)
}

func solve(nums []int, days int) int {
	counts := make([]int, 5)
	for _, num := range nums {
		counts[num-1]++
	}

	resChan := make(chan int, 5)

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		c := counts[i-1]
		if c > 0 {
			wg.Add(1)
			go func(n, count int) {
				defer wg.Done()
				resChan <- calc(i, days) * c
			}(i, c)
		}
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	var res int
	for v := range resChan {
		res += v
	}
	res += len(nums)

	return res
}

func calc(n, days int) int {
	var res int
	if days-n <= 0 {
		return res
	}

	days -= n + 1

	for ; days >= 0; days -= 7 {
		res++
		res += calc(8, days)
	}
	return res
}

func naiveSolve(nums []int, days int) int {
	for range days {
		l := len(nums)
		for i := range l {
			nums[i]--
			if nums[i] == -1 {
				nums[i] = 6
				nums = append(nums, 8)
			}
		}
	}
	return len(nums)
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
