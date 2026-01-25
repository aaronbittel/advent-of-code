package main

import (
	"AOC2022/internal/common"
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type ListType int

const (
	KindNumber ListType = iota
	KindList
)

type List []ListElement

type ListElement struct {
	Kind      ListType
	NumValue  int
	ListValue []ListElement
}

type Pair struct {
	Left  List
	Right List
}

func main() {
	filename := common.GetFilename()

	pairs := parse(filename)

	res1, dur1 := common.TimeIt(func() int {
		return part1(pairs)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)

	res2, dur2 := common.TimeIt(func() int {
		return part2(pairs)
	})
	fmt.Printf("Part2: %d, took %s\n", res2, dur2)
}

func part2(pairs []Pair) int {
	allPackets := make([]List, 0, len(pairs)*2+2)
	for _, pair := range pairs {
		allPackets = append(allPackets, pair.Left, pair.Right)
	}
	allPackets = append(allPackets, ParseList([]byte("[[2]]")))
	allPackets = append(allPackets, ParseList([]byte("[[6]]")))

	slices.SortFunc(allPackets, func(a, b List) int {
		return b.Compare(a)
	})

	res := 1
	for i, list := range allPackets {
		if list.String() == "[[2]]" {
			res *= (i + 1)
		}
		if list.String() == "[[6]]" {
			res *= (i + 1)
		}
	}
	return res
}

func part1(pairs []Pair) int {
	var res int
	for i, pair := range pairs {
		if pair.CorrectOrder() {
			res += i + 1
		}
	}
	return res
}

func (l List) Compare(other List) int {
	minLength := min(len(l), len(other))
	for i := range minLength {
		c := l[i].Compare(other[i])
		if c != 0 {
			return c
		}
	}
	lenDiff := len(l) - len(other)
	if lenDiff < 0 {
		return 1
	} else if lenDiff > 0 {
		return -1
	}
	return 0
}

func (p Pair) CorrectOrder() bool {
	return p.Left.Compare(p.Right) == 1
}

func (le ListElement) Compare(other ListElement) int {
	if le.Kind == KindNumber && other.Kind == KindNumber {
		if le.NumValue < other.NumValue {
			return 1
		} else if le.NumValue > other.NumValue {
			return -1
		}
		return 0
	}
	if le.Kind == KindList && other.Kind == KindList {
		minLength := min(len(le.ListValue), len(other.ListValue))
		for i := range minLength {
			c := le.ListValue[i].Compare(other.ListValue[i])
			if c == 0 {
				continue
			} else {
				return c
			}
		}
		if len(le.ListValue) < len(other.ListValue) {
			return 1
		} else if len(le.ListValue) > len(other.ListValue) {
			return -1
		}
		return 0
	}
	if le.Kind == KindNumber {
		leftAsList := ListElement{
			Kind: KindList,
			ListValue: []ListElement{
				{
					Kind:     KindNumber,
					NumValue: le.NumValue,
				},
			},
		}
		return leftAsList.Compare(other)
	}
	rightAsList := ListElement{
		Kind: KindList,
		ListValue: []ListElement{
			{
				Kind:     KindNumber,
				NumValue: other.NumValue,
			},
		},
	}
	return le.Compare(rightAsList)
}

func parse(filename string) []Pair {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	data = bytes.TrimRight(data, "\n")

	pairs := []Pair{}

	for _, pairRaw := range bytes.Split(data, []byte("\n\n")) {
		parts := bytes.Split(pairRaw, []byte{'\n'})
		leftRaw, rightRaw := parts[0], parts[1]
		left, right := ParseList(leftRaw), ParseList(rightRaw)
		if left.String() != string(leftRaw) {
			fmt.Println("left", string(leftRaw))
			panic("parsing mismatch")
		}
		if right.String() != string(rightRaw) {
			fmt.Println("right", string(rightRaw))
			panic("parsing mismatch")
		}
		pairs = append(pairs, Pair{Left: left, Right: right})
	}
	return pairs
}

type Parser struct {
	Src   []byte
	Index int
}

func ParseList(data []byte) List {
	parser := Parser{Src: data}
	return parser.Parse()
}

func (p *Parser) Parse() List {
	l := List{}
	p.Expect('[')
	for p.Current() != ']' {
		l = append(l, p.ParseElement())
	}
	p.Expect(']')
	return l
}

// [[1],[2,3,4]]
// List[
//
//	ListElement{Kind: List, ListValue: [ListElement{Kind: Num, Value: 1}]},
//	ListElement{Kind: List, ListValue: [
//	  ListElement{Kind: Num, Value: 2},
//	  ListElement{Kind: Num, Value: 3},
//	  ListElement{Kind: Num, Value: 4},
//	],
//
// ]
func (p *Parser) ParseElement() ListElement {
	el := ListElement{Kind: KindList}
	for {
		char := p.Current()
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			start := p.Index
			for p.Current() >= '0' && p.Current() <= '9' {
				p.Index++
			}
			num, err := strconv.Atoi(string(p.Src[start:p.Index]))
			if err != nil {
				log.Fatal(err)
			}
			el = ListElement{Kind: KindNumber, NumValue: num}
			return el
		case ',':
			p.Index++
		case '[':
			p.Expect('[')
			el.Kind = KindList
			for p.Current() != ']' {
				el.ListValue = append(el.ListValue, p.ParseElement())
			}
			p.Expect(']')
			return el
		}
	}
}

func (p *Parser) Expect(char byte) {
	if p.Current() != char {
		panic(fmt.Sprintf("expected %c, but got %c", char, p.Current()))
	}
	p.Index++
}

func (p Parser) Current() byte {
	if p.Index >= len(p.Src) {
		return p.Src[len(p.Src)-1]
	}
	return p.Src[p.Index]
}

func (p Parser) Done() bool {
	return p.Index >= len(p.Src)
}

func (l List) String() string {
	sb := strings.Builder{}
	sb.WriteByte('[')
	for i, item := range l {
		sb.WriteString(item.String())
		if i != len(l)-1 {
			sb.WriteByte(',')
		}
	}
	sb.WriteByte(']')
	return sb.String()
}

func (le ListElement) String() string {
	sb := strings.Builder{}
	if le.Kind == KindNumber {
		sb.WriteString(strconv.Itoa(le.NumValue))
	} else {
		sb.WriteByte('[')
		for i, it := range le.ListValue {
			sb.WriteString(it.String())
			if i != len(le.ListValue)-1 {
				sb.WriteByte(',')
			}
		}
		sb.WriteByte(']')
	}
	return sb.String()
}
