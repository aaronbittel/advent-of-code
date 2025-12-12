package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Machine struct {
	lightTarget    int
	width          int
	buttonMask     []int
	joltageTargets []int
}

type MachinePart2 struct {
	joltageTargets []int
	buttons        [][]int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	machines := parse(f)

	res1 := part1(machines)
	fmt.Printf("Part1: %d\n", res1)

	machinesPart2 := make([]MachinePart2, len(machines))
	for i, m := range machines {
		machinesPart2[i] = FromMachine(m)
	}

	res2 := part2(machinesPart2)
	fmt.Printf("Part2: %d\n", res2)
}

// Shoutout to: https://github.com/michel-kraemer/adventofcode-rust/blob/main/2025/day10/src/main.rs
func part2(machines []MachinePart2) int {
	results := make(chan int, len(machines))

	for _, m := range machines {
		buttonMask := 0
		mask := 1
		for range len(m.buttons) {
			buttonMask |= mask
			mask = mask << 1
		}
		go func(m MachinePart2, buttonMask int) {
			r := joltageButtonPresses(m.joltageTargets, buttonMask, m.buttons)
			results <- r
		}(m, buttonMask)
	}

	var res int
	l := float64(len(machines))

	for i := range len(machines) {
		res += <-results
		fmt.Printf("\r%.2f %% Done.", float64(i+1)/l*100.0)
	}

	return res
}

func nextCombination(combinations []int) bool {
	var idx int
	for i, v := range slices.Backward(combinations) {
		if v != 0 {
			idx = i
			break
		}
	}
	if idx == 0 { // ?
		return false
	}

	v := combinations[idx]
	combinations[idx-1] += 1
	combinations[idx] = 0
	combinations[len(combinations)-1] = v - 1
	return true
}

func isAvailable(mask, i, width int) bool {
	pos := 1 << (width - 1 - i)
	res := (mask & pos)
	return res != 0
}

func joltageButtonPresses(joltages []int, availableButtonMask int, buttons [][]int) int {
	goal := make([]int, len(joltages))
	if slices.Equal(joltages, goal) {
		return 0
	}
	if availableButtonMask == 0 {
		return math.MaxInt
	}
	// find the joltage index that is changed by the least amount of buttons
	indexes := make([]int, len(joltages))
	for i, button := range buttons {
		if !isAvailable(availableButtonMask, i, len(buttons)) {
			continue
		}
		for _, b := range button {
			indexes[b]++
		}
	}

	var (
		buttonIdx  = -1
		minPresses = math.MaxInt
	)

	for i, v := range indexes {
		if v < minPresses && v > 0 {
			minPresses = v
			buttonIdx = i
		}
	}

	type Button struct {
		index   int
		presses []int
	}

	// get the matching buttons
	matchingButtons := make([]Button, 0, minPresses)
	for i, button := range buttons {
		if isAvailable(availableButtonMask, i, len(buttons)) && slices.Contains(button, buttonIdx) {
			matchingButtons = append(matchingButtons, Button{
				index: i, presses: button,
			})
		}
	}
	if len(matchingButtons) == 0 {
		panic("what does this mean?")
	}

	newMask := availableButtonMask
	for _, btn := range matchingButtons {
		pos := len(buttons) - 1 - btn.index
		newMask &= ^(1 << pos)
	}

	result := math.MaxInt

	counts := make([]int, minPresses)
	counts[minPresses-1] = joltages[buttonIdx]
	minP := joltages[buttonIdx]

	for {
		newJoltage := slices.Clone(joltages)
		good := true

	buttons:
		for i, cnt := range counts {
			if cnt == 0 {
				continue
			}
			for _, b := range buttons[matchingButtons[i].index] {
				if newJoltage[b] >= cnt {
					newJoltage[b] -= cnt
				} else {
					good = false
					break buttons
				}
			}

		}

		if good {
			r := joltageButtonPresses(newJoltage, newMask, buttons)
			if r != math.MaxInt {
				result = min(result, minP+r)
			}
		}

		if !nextCombination(counts) {
			break
		}
	}
	return result
}

func part1(machines []Machine) int {
	var res int
	for _, m := range machines {
		res += lightButtonPresses(m)
	}
	return res
}

func lightButtonPresses(m Machine) int {
	type State struct {
		lights  int
		presses int
	}
	queue := []State{{}}
	visited := make(map[int]struct{})
	var minPresses int
	for {
		state := queue[0]
		queue = queue[1:]

		if m.lightTarget == state.lights {
			minPresses = state.presses
			break
		}

		for _, b := range m.buttonMask {
			newLights := pressLightButtons(state.lights, b)
			if _, seen := visited[newLights]; seen {
				continue
			}
			visited[newLights] = struct{}{}
			queue = append(queue, State{lights: newLights, presses: state.presses + 1})
		}
	}
	return minPresses
}

func pressLightButtons(lights, buttonMask int) int {
	return lights ^ buttonMask
}

func parse(f io.Reader) []Machine {
	machines := []Machine{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		m := Machine{}
		line := scanner.Text()
		for _, part := range strings.Fields(line) {
			switch part[0] {
			case '[':
				m.lightTarget, m.width = parseLights(part)
			case '(':
				m.buttonMask = append(m.buttonMask, parseButtons(part, m.width))
			case '{':
				m.joltageTargets = parseJoltage(part)
			default:
				panic("unknown symbol")
			}
		}
		machines = append(machines, m)
	}
	return machines
}

func parseLights(s string) (int, int) {
	width := len(s) - 2
	var val int
	for i, c := range s[1 : len(s)-1] {
		if c == '#' {
			val |= 1 << (width - i - 1)
		}
	}
	return val, width
}

func parseButtons(s string, width int) int {
	if !strings.Contains(s, ",") {
		n, err := strconv.Atoi(string(s[1]))
		if err != nil {
			log.Fatal(err)
		}
		return makeMask(n, width)
	}
	var mask int
	for _, n := range strings.Split(s[1:len(s)-1], ",") {
		n, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		mask |= makeMask(n, width)
	}
	return mask
}

func parseJoltage(s string) []int {
	if !strings.Contains(s, ",") {
		n, err := strconv.Atoi(string(s[1]))
		if err != nil {
			log.Fatal(err)
		}
		return []int{n}
	}

	nums := []int{}
	for _, n := range strings.Split(s[1:len(s)-1], ",") {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func makeMask(n, width int) int {
	return 1 << (width - n - 1)
}

func (m Machine) String() string {
	return fmt.Sprintf("%0*b %v %v", m.width, m.lightTarget, m.buttonMask, m.joltageTargets)
}

func FromMachine(m Machine) MachinePart2 {
	m2 := MachinePart2{
		joltageTargets: m.joltageTargets,
		buttons:        make([][]int, len(m.buttonMask)),
	}

	for bi, bm := range m.buttonMask {
		button := []int{}
		for i := range m.width {
			pos := m.width - i - 1
			if bm&(1<<pos) != 0 {
				button = append(button, i)
			}
		}
		m2.buttons[bi] = button
	}

	return m2
}
