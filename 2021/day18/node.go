package main

import (
	"AOC2021/internal/common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  int
}

func (n *Node) Magnitude() int {
	if n.IsLeaf() {
		return n.Value
	}
	return 3*n.Left.Magnitude() + 2*n.Right.Magnitude()
}

func (n *Node) Sum(other *Node) *Node {
	resultNode := n.Add(other)
	for resultNode.Reduce() {
		common.Assert(resultNode.Validate(), "validate node consistency")
	}
	return resultNode
}

func (n *Node) Reduce() bool {
	if !n.Explode() {
		if !n.Split() {
			return false
		}
	}
	return true
}

func (n *Node) Split() bool {
	if node := n.findSplitNode(); node != nil {
		node.doSplit()
		return true
	}
	return false
}

func (n *Node) doSplit() {
	common.Assert(n.IsLeaf(), "must be a leaf node")
	newNode := &Node{Parent: n.Parent}
	left := &Node{Parent: newNode, Value: n.Value / 2}
	right := &Node{Parent: newNode, Value: int(math.Ceil(float64(n.Value) / 2))}
	newNode.Left = left
	newNode.Right = right
	if n.Parent.Left == n {
		n.Parent.Left = newNode
	} else {
		common.Assert(n.Parent.Right == n, "connection must be correct")
		n.Parent.Right = newNode
	}
}

func (n *Node) findSplitNode() *Node {
	if n.IsLeaf() {
		if n.Value >= 10 {
			return n
		}
		return nil
	}
	if node := n.Left.findSplitNode(); node != nil {
		return node
	}
	if node := n.Right.findSplitNode(); node != nil {
		return node
	}
	return nil
}

func (n *Node) doExplode() {
	common.Assert(n.Left.IsLeaf(), "error: left node must be a leaf")

	leftValue := n.Left.Value
	last := n
	var updatedLeft bool
	for cur := n.Parent; cur.Parent != nil; cur = cur.Parent {
		if cur.Left != nil && cur.Left != last {
			cur = cur.Left
			for cur.Right != nil {
				cur = cur.Right
			}
			updatedLeft = true
			cur.Value += leftValue
			break
		}
		last = cur
	}

	// after reaching the root without finding a leaf to update
	// search left from the root on the right side
	// to find the next value to the left from the target node
	// only if original node was on the right side of the tree
	if !n.IsLeftSide() && !updatedLeft {
		for cur := n.Root().Left; cur != nil; cur = cur.Right {
			if cur.Right == nil {
				cur.Value += leftValue
				break
			}
		}
	}

	last = n
	rightValue := n.Right.Value
	var updatedRight bool
	for cur := n.Parent; cur.Parent != nil; cur = cur.Parent {
		if cur.Right != nil && cur.Right != last {
			cur = cur.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Value += rightValue
			updatedRight = true
			break
		}
		last = cur
	}

	// after reaching the root without finding a leaf to update
	// search right from the root on the left side
	// to find the next value to the right from the target node
	// only if original node was on the left side of the tree
	if n.IsLeftSide() && !updatedRight {
		for cur := n.Root().Right; cur != nil; cur = cur.Left {
			if cur.Left == nil {
				cur.Value += rightValue
				break
			}
		}
	}

	newNode := &Node{Parent: n.Parent}
	if n.Parent.Left == n {
		n.Parent.Left = newNode
	} else {
		common.Assert(n.Parent.Right == n, "connection must be correct")
		n.Parent.Right = newNode
	}

}

func (n *Node) Explode() bool {
	if node := n.findExplodePair(0); node != nil {
		node.doExplode()
		return true
	}
	return false
}

func (n *Node) findExplodePair(level int) *Node {
	if level >= 4 && n.Left.IsLeaf() && n.Right.IsLeaf() {
		return n
	}
	if !n.Left.IsLeaf() {
		if node := n.Left.findExplodePair(level + 1); node != nil {
			return node
		}
	}
	if !n.Right.IsLeaf() {
		if node := n.Right.findExplodePair(level + 1); node != nil {
			return node
		}
	}
	return nil
}

func (n *Node) Add(other *Node) *Node {
	newNode := &Node{Left: n, Right: other}
	n.Parent = newNode
	other.Parent = newNode
	return newNode
}

func ParseNode(s string) *Node {
	p := Parser{Src: s}
	return p.Parse(nil)
}

func (n *Node) String() string {
	sb := strings.Builder{}
	if n.IsLeaf() {
		sb.WriteString(strconv.Itoa(n.Value))
	} else {
		sb.WriteString(fmt.Sprintf("[%s,%s]", n.Left, n.Right))
	}
	return sb.String()
}

func (n Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func (p *Parser) Expect(ch byte) {
	if p.Src[p.Idx] != ch {
		panic(fmt.Sprintf("error: expected %s, but got %s", string(ch), string(p.Char())))
	}
	p.Idx++
}

func (p Parser) Char() byte {
	return p.Src[p.Idx]
}

func (n *Node) Validate() bool {
	if n.Left.Parent != n {
		return false
	}
	if !n.Left.IsLeaf() && !n.Left.Validate() {
		return false
	}
	if n.Right.Parent != n {
		return false
	}
	if !n.Right.IsLeaf() && !n.Right.Validate() {
		return false
	}
	return true
}

func (n *Node) Root() *Node {
	root := n
	for root.Parent != nil {
		root = root.Parent
	}
	return root
}

func (n *Node) IsLeftSide() bool {
	root := n.Root()
	if n == root {
		panic("illegal node")
	}
	last := n
	for last.Parent.Parent != nil {
		last = last.Parent
	}
	return root.Left == last
}

func (n *Node) Clone() *Node {
	return ParseNode(n.String())
}
