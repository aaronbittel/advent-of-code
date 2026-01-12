package main

import (
	"AOC2021/internal/common"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Parser struct {
	data string
	idx  int

	totalNumVersion int
}

func main() {
	filename := common.GetFilename()

	packet := parse(filename)
	data := decodeToBinary(packet)

	res1, dur1 := common.TimeIt(func() int {
		return part1(data)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)
}

func part1(data string) int {
	parser := Parser{data: data}
	parser.packet()

	return parser.totalNumVersion
}

func (p *Parser) packet() int {
	var read int

	version := p.version()
	read += 3
	p.totalNumVersion += version
	typeID := p.typeID()
	read += 3
	if typeID == 4 {
		_, readLit := p.literal()
		read += readLit
	} else {
		readOp := p.operator()
		read += readOp
	}
	return read
}

func (p *Parser) operator() int {
	var read int
	if p.parseN(1) == 0 {
		totalLength := p.parseN(15)
		read += 16
		for totalLength > 0 {
			packetRead := p.packet()
			read += packetRead
			totalLength -= packetRead
		}
	} else {
		numSubpackets := p.parseN(11)
		read += 12
		for range numSubpackets {
			read += p.packet()
		}
	}
	return read
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
	switch hex {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'A':
		return "1010"
	case 'B':
		return "1011"
	case 'C':
		return "1100"
	case 'D':
		return "1101"
	case 'E':
		return "1110"
	case 'F':
		return "1111"
	default:
		panic("invalid hex string")
	}
}
