package main

import (
	"log"
	"strconv"
)

type Parser struct {
	Src string
	Idx int
}

func (p *Parser) Parse(parent *Node) *Node {
	ch := p.Src[p.Idx]
	if ch >= '0' && ch <= '9' {
		start := p.Idx
		p.Idx++ // consume digit
		for p.Char() >= '0' && p.Char() <= '9' {
			p.Idx++
		}
		n, err := strconv.Atoi(p.Src[start:p.Idx])
		if err != nil {
			log.Fatal(err)
		}
		return &Node{Parent: parent, Value: n}
	}
	newNode := &Node{Parent: parent}
	p.Expect('[')
	newNode.Left = p.Parse(newNode)
	p.Expect(',')
	newNode.Right = p.Parse(newNode)
	p.Expect(']')
	return newNode
}
