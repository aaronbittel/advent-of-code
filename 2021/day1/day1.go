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

	var (
		part1 int
		prev  = math.MaxInt

		part2      int
		prevWindow = math.MaxInt
		window     = [3]int{}
		windowSum  int
	)

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		// part1
		if num > prev {
			part1++
		}
		prev = num

		// part2
		windowSum += num
		if i >= 3 {
			windowSum -= window[i%3]
			if windowSum > prevWindow {
				part2++
			}
		}
		window[i%3] = num
		prevWindow = windowSum
	}

	fmt.Printf("Par1: %d\n", part1)
	fmt.Printf("Par2: %d\n", part2)
}
