package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	nums := parse(f)

	res1 := part1(nums)
	fmt.Printf("Part1: %d\n", res1)

	res2 := part2(nums)
	fmt.Printf("Part2: %d\n", res2)
}

func part2(nums []int) int {
	const COUNT = 2000
	memos := make([]map[[4]int]int, 0)

	for _, num := range nums {
		secretNums := make([]int, COUNT)
		for i := range COUNT {
			secretNums[i] = num % 10
			num = nextSecret(num)
		}

		changes := make([]int, 0, COUNT-1)
		for i := 1; i < len(secretNums); i++ {
			changes = append(changes, secretNums[i]-secretNums[i-1])
		}

		memo := make(map[[4]int]int)
		r := 4
		l := 0
		for r < len(changes) {
			window := [4]int{changes[l], changes[l+1], changes[l+2], changes[l+3]}
			if _, ok := memo[window]; !ok {
				memo[window] = secretNums[r]
			}
			l++
			r++
		}

		memos = append(memos, memo)
	}

	for _, memo := range memos[1:] {
		for k, v := range memo {
			memos[0][k] += v
		}
	}

	var res int
	for _, v := range memos[0] {
		if v > res {
			res = v
		}
	}
	return res
}

func part1(nums []int) int {
	var res int
	for _, n := range nums {
		secret := n
		for range 2000 {
			secret = nextSecret(secret)
		}
		res += secret
	}
	return res
}

func nextSecret(s int) int {
	newS := s
	newS = mix(newS, newS*64)
	newS = prune(newS)
	newS = mix(newS, int(math.Floor(float64(newS)/32.0)))
	newS = prune(newS)
	newS = mix(newS, newS*2048)
	newS = prune(newS)
	return newS
}

func mix(s, v int) int {
	return s ^ v
}

func prune(s int) int {
	return s % 16777216
}

func parse(f io.Reader) []int {
	nums := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, n)
	}
	return nums
}
