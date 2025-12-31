package main

import (
	"AOC2021/internal/common"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile(common.GetFilename())
	if err != nil {
		log.Fatal(err)
	}
	content := string(data)
	report := parse(content)

	part1, dur1 := common.TimeIt(func() int {
		return part1(report)
	})
	fmt.Printf("Part1: %d, took %s\n", part1, dur1)

	part2, dur2 := common.TimeIt(func() int {
		return part2(report)
	})
	fmt.Printf("Part2: %d, took %s\n", part2, dur2)
}

func part2(report []string) int {
	oxygenReport := make([]string, len(report))
	copy(oxygenReport, report)
	for i := 0; len(oxygenReport) != 1; i++ {
		oxygenReport = oxygenRating(split(oxygenReport, i))
	}
	oxygen, err := strconv.ParseInt(oxygenReport[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	co2Report := make([]string, len(report))
	copy(co2Report, report)
	for i := 0; len(co2Report) != 1; i++ {
		co2Report = co2Rating(split(co2Report, i))
	}
	co2, err := strconv.ParseInt(co2Report[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return int(oxygen) * int(co2)
}

func split(report []string, index int) ([]string, []string) {
	ones := make([]string, 0, len(report))
	zeros := make([]string, 0, len(report))

	for _, line := range report {
		if line[index] == '1' {
			ones = append(ones, line)
		} else {
			zeros = append(zeros, line)
		}
	}

	return ones, zeros
}

func oxygenRating(ones, zeros []string) []string {
	if len(ones) >= len(zeros) {
		return ones
	}
	return zeros
}

func co2Rating(ones, zeros []string) []string {
	if len(zeros) <= len(ones) {
		return zeros
	}
	return ones
}

func part1(report []string) int {
	counter := make([]int, len(report[0]))

	for _, line := range report {
		for i, r := range slices.Backward([]byte(line)) {
			if r == '1' {
				counter[i]++
			}
		}
	}

	var (
		gammaRate   int
		epsilonRate int
		tmp         int = 1
	)

	halfLineCount := len(report) / 2
	for _, c := range slices.Backward(counter) {
		if c > halfLineCount {
			gammaRate += tmp
		} else {
			epsilonRate += tmp
		}
		tmp <<= 1
	}

	return gammaRate * epsilonRate
}

func parse(content string) []string {
	lines := strings.Split(content, "\n")
	last := len(lines) - 1
	return lines[:last]
}
