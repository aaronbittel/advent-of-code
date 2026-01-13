package main

import (
	"AOC2021/internal/common"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	data string
	idx  int

	totalNumVersion int
}

type Operator int

const (
	Sum Operator = iota
	Product
	Minimum
	Maximum
	Literal
	GreaterThan
	LessThan
	EqualTo
)

func main() {
	filename := common.GetFilename()

	packet := parse(filename)
	data := decodeToBinary(packet)

	res, dur1 := common.TimeIt(func() [2]int {
		res1, res2 := part1(data)
		return [2]int{res1, res2}
	})
	fmt.Printf("Part1: %d, Part2: %d, took %s\n", res[0], res[1], dur1)
}

func part1(data string) (int, int) {
	parser := Parser{data: data}
	res, _ := parser.packet()

	return parser.totalNumVersion, res
}

func (p *Parser) packet() (int, int) {
	var (
		read int
		res  int
	)

	version := p.version()
	read += 3
	p.totalNumVersion += version
	op := Operator(p.typeID())
	read += 3
	if op == Literal {
		val, readLit := p.literal()
		read += readLit
		res = val
	} else {
		val, readOp := p.operator(op)
		read += readOp
		res = val
	}
	return res, read
}

func (p *Parser) operator(op Operator) (int, int) {
	var (
		read int
		res  int

		cmp1 int
		cmp2 int
	)
	switch op {
	case Sum:
		res = 0
	case Product:
		res = 1
	case Minimum:
		res = math.MaxInt
	case Maximum:
		res = math.MinInt
	case GreaterThan, LessThan, EqualTo:
		cmp1, cmp2 = -1, -1
	default:
		panic("illegal operator")
	}
	if p.parseN(1) == 0 {
		totalLength := p.parseN(15)
		read += 16
		for totalLength > 0 {
			val, packetRead := p.packet()
			read += packetRead
			totalLength -= packetRead
			if op == GreaterThan || op == LessThan || op == EqualTo {
				if cmp1 == -1 {
					cmp1 = val
				} else if cmp2 == -1 {
					cmp2 = val
					res = apply(op, cmp1, cmp2)
				}
			} else {
				res = apply(op, res, val)
			}
		}
	} else {
		numSubpackets := p.parseN(11)
		read += 12
		for range numSubpackets {
			val, packetRead := p.packet()
			read += packetRead
			if op == GreaterThan || op == LessThan || op == EqualTo {
				if cmp1 == -1 {
					cmp1 = val
				} else if cmp2 == -1 {
					cmp2 = val
					res = apply(op, cmp1, cmp2)
				}
			} else {
				res = apply(op, res, val)
			}
		}
	}
	return res, read
}

func (p *Parser) literal() (int, int) {
	var (
		numStr string
		read   int
	)

	for p.parseN(1) == 1 {
		numStr += p.readN(4)
		read += 5
	}
	numStr += p.readN(4)
	read += 5

	num64, err := strconv.ParseInt(numStr, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(num64), read
}

func (p *Parser) version() int {
	return p.parseN(3)
}

func (p *Parser) typeID() int {
	return p.parseN(3)
}

func (p *Parser) readN(n int) string {
	data := p.data[p.idx : p.idx+n]
	p.idx += n
	return data
}

func (p *Parser) parseN(n int) int {
	i64, err := strconv.ParseInt(p.readN(n), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(i64)
}

func decodeToBinary(raw string) string {
	sb := strings.Builder{}
	sb.Grow(len(raw) * 4)
	for i := range raw {
		sb.WriteString(hexToBinStr(raw[i]))
	}
	return sb.String()
}

func parse(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(data[:len(data)-1])
}

func hexToBinStr(hex byte) string {
	i64, err := strconv.ParseUint(string(hex), 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%04b", i64)
}

func apply(op Operator, v1, v2 int) int {
	switch op {
	case Sum:
		return v1 + v2
	case Product:
		return v1 * v2
	case Minimum:
		return min(v1, v2)
	case Maximum:
		return max(v1, v2)
	case GreaterThan:
		if v1 > v2 {
			return 1
		}
		return 0
	case LessThan:
		if v1 < v2 {
			return 1
		}
		return 0
	case EqualTo:
		if v1 == v2 {
			return 1
		}
		return 0
	default:
		panic("illegal operator")
	}
}

func (op Operator) String() string {
	switch op {
	case Sum:
		return "Sum"
	case Product:
		return "Product"
	case Minimum:
		return "Minimum"
	case Maximum:
		return "Maximum"
	case Literal:
		return "Literal"
	case GreaterThan:
		return "GreaterThan"
	case LessThan:
		return "LessThan"
	case EqualTo:
		return "EqualTo"
	default:
		panic("illegal operator")
	}
}
