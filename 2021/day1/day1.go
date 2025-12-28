package main

import (
	"bufio"
	"fmt"
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
	defer f.Close()

	prev1 := math.MaxInt
	prev2 := math.MaxInt
	window := []int{}
	var res1 int
	var res2 int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if num > prev1 {
			res1++
		}
		prev1 = num

		if len(window) < 3 {
			window = append(window, num)
		}

		if len(window) == 3 {
			if Sum(window) > prev2 {
				res2++
			}
			prev2 = Sum(window)
			window = window[1:]
		}

	}

	fmt.Printf("Par1: %d\n", res1)
	fmt.Printf("Par2: %d\n", res2)
}

func Sum(s []int) int {
	var res int
	for _, xs := range s {
		res += xs
	}
	return res
}
