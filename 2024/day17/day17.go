package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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

	vm := parse(f)

	res1 := part1(vm)
	fmt.Printf("Part1: %s\n", res1)
}

func part1(vm VM) string {
	for !vm.AtEnd() {
		vm.Step()
	}

	return vm.Result()
}

func parse(f io.Reader) VM {
	scanner := bufio.NewScanner(f)
	regs := [3]int{}
	for i := 0; scanner.Scan() && i < 3; i++ {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, ": ")
		reg, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		regs[i] = reg
	}

	if !scanner.Scan() {
		log.Fatal("no program provided")
	}

	programLine := scanner.Text()
	parts := strings.Split(programLine, ": ")
	program := []int{}

	for _, rp := range strings.Split(parts[1], ",") {
		p, err := strconv.Atoi(rp)
		if err != nil {
			log.Fatal(err)
		}
		program = append(program, p)
	}

	return NewVM(regs, program)
}
