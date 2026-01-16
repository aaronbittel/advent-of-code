package main

import (
	"AOC2021/internal/common"
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Display struct {
	Patterns []string
	Output   []string
}

type Segment struct {
	Top, TopLeft, TopRight, Middle, BottomLeft, BottomRight, Bottom byte
}

func main() {
	f := common.GetFile()
	defer f.Close()

	displays := parse(f)

	res1, dur1 := common.TimeIt(func() int {
		return part1(displays)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)

	res2, dur2 := common.TimeIt(func() int {
		return part2(displays)
	})
	fmt.Printf("Part2: %d, took %s\n", res2, dur2)
}

func part2(displays []Display) int {
	var res int
	for _, d := range displays {
		res += d.Result()
	}
	return res
}

func part1(displays []Display) int {
	var res int
	for _, d := range displays {
		for _, o := range d.Output {
			if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
				res++
			}
		}
	}
	return res
}

func (d Display) Result() int {
	s := Segment{}
	combined := d.Combined()

	ones := Map(combined, func(s string) bool { return len(s) == 2 })
	fours := []string{}
	if len(ones) > 0 {
		fours = Map(combined, func(s string) bool { return len(s) == 4 })
		if len(fours) > 0 {
			dis, _ := disjunct(ones[0], fours[0])
			common.Assert(len(dis) == 2, "expected 2")
			for _, c := range combined {
				if len(c) != 5 {
					continue
				}
				if strings.Contains(c, string(dis[0])) && strings.Contains(c, string(dis[1])) {
					_, commonChar := disjunct(c, ones[0])
					common.Assert(len(commonChar) == 1, "expected one common segment")
					s.BottomRight = []byte(commonChar)[0]
					unique, _ := disjunct(ones[0], commonChar)
					s.TopRight = []byte(unique)[0]
					break
				}
			}
		} else {
			panic("4")
		}
	} else {
		panic("1")
	}

	seven := Map(combined, func(s string) bool { return len(s) == 3 })[0]
	unique, _ := disjunct(ones[0], seven)

	s.Top = []byte(unique)[0]

	var three string
	for _, c := range combined {
		if len(c) != 5 {
			continue
		}
		if strings.Contains(c, string(string(seven[0]))) && strings.Contains(c, string(string(seven[1]))) && strings.Contains(c, string(seven[2])) {
			three = c
			break
		}
	}

	unique3, _ := disjunct(three, seven)
	_, common4 := disjunct(unique3, fours[0])
	s.Middle = []byte(common4)[0]
	u3, _ := disjunct(unique3, common4)
	s.Bottom = []byte(u3)[0]

	for i := range fours[0] {
		if fours[0][i] != s.TopRight && fours[0][i] != s.BottomRight && fours[0][i] != s.Middle {
			s.TopLeft = fours[0][i]
		}
	}

	for _, ch := range []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'} {
		if ch != s.Top && ch != s.TopRight && ch != s.TopLeft && ch != s.Middle && ch != s.BottomRight && ch != s.Bottom {
			s.BottomLeft = ch
		}
	}

	var num int
	for _, o := range d.Output {
		var n int
		switch len(o) {
		case 2:
			n = 1
		case 3:
			n = 7
		case 4:
			n = 4
		case 7:
			n = 8
		case 5:
			if strings.Contains(o, string(s.TopLeft)) {
				n = 5
				break
			}
			if strings.Contains(o, string(s.BottomLeft)) {
				n = 2
				break
			}
			n = 3
		case 6:
			if !strings.Contains(o, string(s.Middle)) {
				n = 0
				break
			}
			if strings.Contains(o, string(s.BottomLeft)) {
				n = 6
				break
			}
			n = 9
		default:
			panic("illegal output")
		}
		num = num*10 + n
	}

	return num
}

func disjunct(as, bs string) (string, string) {
	var unique, common string
	if len(bs) > len(as) {
		as, bs = bs, as
	}
	for _, a := range as {
		if strings.Contains(bs, string(a)) {
			common += string(a)
		} else {
			unique += string(a)
		}
	}
	return unique, common
}

func Map(xs []string, fn func(s string) bool) []string {
	out := []string{}
	for _, x := range xs {
		if fn(x) {
			out = append(out, x)
		}
	}
	return out
}

func (d Display) Combined() []string {
	return append(d.Patterns, d.Output...)
}

func parse(r io.Reader) []Display {
	displays := []Display{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " | ")
		patterns := strings.Fields(parts[0])
		output := strings.Fields(parts[1])
		common.Assert(len(output) == 4, "output expected to be of length 4")
		displays = append(displays, Display{
			Patterns: patterns,
			Output:   output,
		})
	}
	return displays
}

func (s Segment) String() string {
	convert := func(b byte) byte {
		if b != 0 {
			return b
		}
		return '.'
	}

	top := convert(s.Top)
	topLeft := convert(s.TopLeft)
	topRight := convert(s.TopRight)
	middle := convert(s.Middle)
	bottomLeft := convert(s.BottomLeft)
	bottomRight := convert(s.BottomRight)
	bottom := convert(s.Bottom)

	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(" %c%c%c%c \n", top, top, top, top))
	sb.WriteString(fmt.Sprintf("%c    %c\n", topLeft, topRight))
	sb.WriteString(fmt.Sprintf("%c    %c\n", topLeft, topRight))
	sb.WriteString(fmt.Sprintf(" %c%c%c%c \n", middle, middle, middle, middle))
	sb.WriteString(fmt.Sprintf("%c    %c\n", bottomLeft, bottomRight))
	sb.WriteString(fmt.Sprintf("%c    %c\n", bottomLeft, bottomRight))
	sb.WriteString(fmt.Sprintf(" %c%c%c%c \n", bottom, bottom, bottom, bottom))
	return sb.String()
}
