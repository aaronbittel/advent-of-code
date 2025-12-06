package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	stones := parse(content)
	res1 := simulate(stones, 25)
	fmt.Printf("\rPart1: %d\n", res1)

	res2 := predict(stones, 75)
	fmt.Printf("Part2: %d\n", res2)
}

type Key struct {
	num        int
	blinksLeft int
}

func predict(stones []int, blinks int) int {
	var res int

	memo := make(map[Key]int)
	for _, num := range stones {
		res += countStones(num, blinks, memo)
	}

	return res
}

func countStones(num int, blinksLeft int, memo map[Key]int) int {
	if blinksLeft == 0 {
		return 1
	}

	key := Key{num, blinksLeft}
	if v, ok := memo[key]; ok {
		return v
	}

	var res int

	switch {
	case num == 0:
		res = countStones(1, blinksLeft-1, memo)

	case evenDigits(num):
		left, right := split(num)
		res = countStones(left, blinksLeft-1, memo) +
			countStones(right, blinksLeft-1, memo)

	default:
		res = countStones(num*2024, blinksLeft-1, memo)
	}

	memo[key] = res
	return res
}

func simulate(stones []int, blinks int) int {
	for range blinks {
		stones = blink(stones)
	}
	return len(stones)
}

func blink(stones []int) []int {
	evolvedStones := make([]int, 0, len(stones))

	for _, stone := range stones {
		switch {
		case stone == 0:
			evolvedStones = append(evolvedStones, 1)
		case evenDigits(stone):
			left, right := split(stone)
			evolvedStones = append(evolvedStones, left, right)
		default:
			evolvedStones = append(evolvedStones, stone*2024)
		}
	}
	return evolvedStones
}

func parse(content string) []int {
	parts := strings.Fields(content)
	stones := make([]int, len(parts))
	for i, part := range parts {
		stones[i] = Atoi(part)
	}
	return stones
}

func evenDigits(x int) bool {
	return len(strconv.Itoa(x))%2 == 0
}

func split(x int) (int, int) {
	s := strconv.Itoa(x)
	mid := len(s) / 2
	left := Atoi(s[:mid])
	right := Atoi(s[mid:])
	return left, right
}

func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}
