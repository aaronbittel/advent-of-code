package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Opcode uint8

const (
	Adv Opcode = iota
	Bxl
	Bst
	Jnz
	Bxc
	Out
	Bdv
	Cdv
)

type Instruction struct {
	Code    Opcode
	Operand int
}

type VM struct {
	Program []Instruction
	A, B, C int

	ip int

	Out []int

	RawProg []int
}

func NewVM(regs [3]int, program []int) VM {
	vm := VM{}
	vm.A = regs[0]
	vm.B = regs[1]
	vm.C = regs[2]

	vm.Program = make([]Instruction, 0, len(program)/2)

	for i := 0; i < len(program); i += 2 {
		vm.Program = append(vm.Program, Instruction{
			Code:    Opcode(program[i]),
			Operand: program[i+1],
		})
	}

	vm.Out = []int{}
	vm.RawProg = program

	return vm
}

func (vm *VM) Run() {
	for !vm.AtEnd() {
		vm.Step()
	}
}

func (vm *VM) Step() {
	inst := vm.Program[vm.ip]
	vm.ip++
	switch inst.Code {
	case Adv:
		numinator := vm.A
		denominator := int(math.Pow(2.0, float64(vm.comboValue(inst.Operand))))
		vm.A = numinator / denominator
	case Bxl:
		vm.B ^= inst.Operand
	case Bst:
		vm.B = vm.comboValue(inst.Operand) % 8
	case Jnz:
		if vm.A == 0 {
			break
		}
		vm.ip = inst.Operand / 2 // Opcode + Operand combined into Instruction, so ip only points at the next Instruction
	case Bxc:
		vm.B ^= vm.C
	case Out:
		v := vm.comboValue(inst.Operand) % 8
		vm.Out = append(vm.Out, v)
	case Bdv:
		numinator := vm.A
		denominator := int(math.Pow(2.0, float64(vm.comboValue(inst.Operand))))
		vm.B = numinator / denominator
	case Cdv:
		numinator := vm.A
		denominator := int(math.Pow(2.0, float64(vm.comboValue(inst.Operand))))
		vm.C = numinator / denominator
	default:
		panic("unknown opcode")
	}
}

func (vm VM) comboValue(n int) int {
	switch n {
	case 0, 1, 2, 3:
		return n
	case 4:
		return vm.A
	case 5:
		return vm.B
	case 6:
		return vm.C
	default:
		panic("invalid operand")
	}
}

func (vm VM) AtEnd() bool {
	return len(vm.Program) <= vm.ip
}

func (vm VM) Result() string {
	b := strings.Builder{}

	for _, o := range vm.Out {
		b.WriteString(fmt.Sprintf("%d,", o))
	}

	out := b.String()
	if len(out) > 0 {
		out = out[:len(out)-1] // remove last ',' again
	}
	return out
}

func (vm VM) String() string {
	b := strings.Builder{}

	b.WriteString(fmt.Sprintf("Register A: %d\n", vm.A))
	b.WriteString(fmt.Sprintf("Register B: %d\n", vm.B))
	b.WriteString(fmt.Sprintf("Register C: %d\n", vm.C))

	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Program: %v\n", vm.Program[vm.ip:]))

	return b.String()
}

func (op Opcode) String() string {
	switch op {
	case Adv:
		return "Adv"
	case Bst:
		return "Bst"
	case Bxl:
		return "Bxl"
	case Jnz:
		return "Jnz"
	case Bxc:
		return "Bxc"
	case Out:
		return "Out"
	case Bdv:
		return "Bdv"
	case Cdv:
		return "Cdv"
	default:
		panic("unknown opcode")
	}
}

func (in Instruction) String() string {
	var opRepr string
	switch in.Operand {
	case 0, 1, 2, 3, 7:
		opRepr = strconv.Itoa(in.Operand)
	case 4:
		opRepr = "A"
	case 5:
		opRepr = "B"
	case 6:
		opRepr = "C"
	default:
		panic("invalid program")
	}

	return fmt.Sprintf("%s(%s)", in.Code, opRepr)
}
