package main

import (
	"AOC2021/internal/common"
	"bufio"
	"fmt"
	"io"
	"log"
	"maps"
	"math"
	"strings"
)

type LinkedList struct {
	Head *Node
	Tail *Node

	Counts map[byte]int
}

type Node struct {
	Next  *Node
	Value byte
}

type Rules map[string]byte

func main() {
	f := common.GetFile()
	defer f.Close()

	template, rules := parse(f)

	res1, dur1 := common.TimeIt(func() int {
		return part1(template, rules, 10)
	})
	fmt.Printf("Part 1: %d, took %s\n", res1, dur1)
}

func part1(template *LinkedList, rules Rules, count int) int {
	for range count {
		step(template, rules)
	}
	var (
		maxCount = math.MinInt
		minCount = math.MaxInt
	)
	for count := range maps.Values(template.Counts) {
		maxCount = max(maxCount, count)
		minCount = min(minCount, count)
	}
	return maxCount - minCount
}

func step(template *LinkedList, rules Rules) {
	cur := template.Head
	next := cur.Next
	for cur != nil && next != nil {
		rule := fmt.Sprintf("%c%c", cur.Value, next.Value)
		if v, ok := rules[rule]; ok {
			template.Counts[v]++

			// Insert new node inbetween cur and next
			newNode := NewNode(v, next)
			cur.Next = newNode
		}

		cur = next
		next = next.Next
	}
}

func parse(r io.Reader) (*LinkedList, Rules) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		panic("invalid input")
	}
	ll := NewLinkedList(scanner.Text())

	if !scanner.Scan() {
		panic("invalid input")
	}
	rules := Rules{}
	for scanner.Scan() {
		var (
			from string
			to   byte
		)
		n, err := fmt.Sscanf(scanner.Text(), "%s -> %c", &from, &to)
		if err != nil {
			log.Fatal(err)
		}
		if n != 2 {
			log.Fatal("expected 2 values")
		}
		rules[from] = to
	}
	return ll, rules
}

func NewLinkedList(template string) *LinkedList {
	ll := LinkedList{
		Counts: make(map[byte]int),
	}
	for i := range template {
		char := template[i]
		newNode := NewNode(char)
		ll.Counts[char]++
		if ll.Head == nil {
			ll.Head = newNode
			ll.Tail = newNode
		} else {
			ll.Tail.Next = newNode
			ll.Tail = newNode
		}
	}
	return &ll
}

func (ll LinkedList) String() string {
	sb := strings.Builder{}

	cur := ll.Head
	for cur != nil {
		fmt.Fprintf(&sb, "%c", cur.Value)
		cur = cur.Next
		if cur != nil {
			sb.WriteString(" -> ")
		}
	}

	return sb.String()
}

func (ll LinkedList) Result() string {
	sb := strings.Builder{}

	cur := ll.Head
	for cur != nil {
		fmt.Fprintf(&sb, "%c", cur.Value)
		cur = cur.Next
	}

	return sb.String()
}

// NewNode returns a new Node with the given value and an optional next node.
func NewNode(value byte, next ...*Node) *Node {
	if len(next) > 1 {
		panic("expected at most one next node")
	}
	var n *Node
	if len(next) == 1 {
		n = next[0]
	}
	return &Node{
		Value: value,
		Next:  n,
	}
}
