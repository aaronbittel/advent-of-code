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

const (
	BOARD_SIZE = 5
	MARKED     = -1
)

type Board [BOARD_SIZE][BOARD_SIZE]int

func main() {
	filename := common.GetFilename()
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	boards, nums := parse(string(data))

	part1, dur1 := common.TimeIt(func() int {
		part1Boards := slices.Clone(boards)
		return part1(part1Boards, nums)
	})
	fmt.Printf("Part1: %d, took %s\n", part1, dur1)

	part2, dur2 := common.TimeIt(func() int {
		return part2(boards, nums)
	})
	fmt.Printf("Part2: %d, took %s\n", part2, dur2)
}

func part2(boards []Board, nums []int) int {
	var (
		losingBoard Board
		losingNum   int
	)

	for _, num := range nums {
		toBeRemoved := []int{}
		for i := range boards {
			if boards[i].Mark(num) && boards[i].Check() {
				toBeRemoved = append(toBeRemoved, i)
			}
		}
		if len(toBeRemoved) == 0 {
			continue
		}
		if len(boards) == 1 {
			losingBoard = boards[0]
			losingNum = num
			break
		}

		boards = filter(boards, toBeRemoved)
	}

	return losingBoard.SumUnmarked() * losingNum
}

func filter(boards []Board, toBeRemoved []int) []Board {
	dst := boards[:0]
	j := 0
	for i := range boards {
		if j < len(toBeRemoved) && i == toBeRemoved[j] {
			j++
			continue
		}
		dst = append(dst, boards[i])
	}
	return dst
}

func part1(boards []Board, nums []int) int {
	var (
		winnerBoard Board
		winnerNum   int
	)
outer:
	for _, num := range nums {
		for i := range boards {
			if boards[i].Mark(num) && boards[i].Check() {
				winnerBoard = boards[i]
				winnerNum = num
				break outer
			}
		}
	}

	return winnerBoard.SumUnmarked() * winnerNum
}

func (b Board) SumUnmarked() int {
	var sum int
	for row := range BOARD_SIZE {
		for col := range BOARD_SIZE {
			if b[row][col] != MARKED {
				sum += b[row][col]
			}
		}
	}
	return sum
}

func (b Board) Check() bool {
	for _, row := range b {
		won := true
		for _, n := range row {
			if n != MARKED {
				won = false
				break
			}
		}
		if won {
			return true
		}
	}

	for col := range BOARD_SIZE {
		won := true
		for row := range BOARD_SIZE {
			if b[row][col] != MARKED {
				won = false
				break
			}
		}
		if won {
			return true
		}
	}

	return false
}

func (b *Board) Mark(num int) bool {
	var updated bool
	for row := range BOARD_SIZE {
		for col := range BOARD_SIZE {
			if b[row][col] == num {
				b[row][col] = MARKED
				updated = true
			}
		}
	}
	return updated
}

func parse(content string) ([]Board, []int) {
	var (
		boards = []Board{}
		nums   = []int{}
	)

	lines := strings.Split(content, "\n")
	last := len(lines) - 1
	numLine, boardLines := lines[0], lines[2:last]

	// parse nums
	for _, numStr := range strings.Split(numLine, ",") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	//parse boards
	for i := 0; i < len(boardLines); i += 6 {
		var board Board
		start := i
		for j := start; j < start+BOARD_SIZE; j++ {
			for col, numStr := range strings.Fields(boardLines[j]) {
				num, err := strconv.Atoi(numStr)
				if err != nil {
					log.Fatal(err)
				}
				board[j-i][col] = num
			}
		}
		boards = append(boards, board)
	}

	return boards, nums
}

func (b Board) String() string {
	sb := strings.Builder{}
	for row := range BOARD_SIZE {
		for col := range BOARD_SIZE {
			fmt.Fprintf(&sb, "%2d ", b[row][col])
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
