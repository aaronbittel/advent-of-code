package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "USAGE: %s <input-file>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	part1 := part2(f, 2)

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	part2 := part2(f, 12)

	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}

func part2(f io.Reader, count int) int {
	scanner := bufio.NewScanner(f)
	var res int
	for scanner.Scan() {
		line := scanner.Text()
		start := 0
		end := len(line) - count + 1
		var r int
		for range count {
			num, idx := find_highest_joltage(line[start:end])
			start += idx + 1
			end += 1
			r = r*10 + num
		}
		res += r
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}

func find_highest_joltage(line string) (int, int) {
	var num, idx int
	for i, char := range line {
		n, err := strconv.Atoi(string(char))
		if err != nil {
			log.Fatal(err)
		}
		if n > num {
			num = n
			idx = i
			if num == 9 {
				break
			}
		}
	}
	return num, idx
}
