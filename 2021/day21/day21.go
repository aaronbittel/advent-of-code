package main

import (
	"AOC2021/internal/common"
	"bufio"
	"fmt"
	"io"
	"log"
)

type Dice struct {
	NextValue int
	Rolled    int
}

type Player struct {
	Pos   int
	Score int
}

func main() {
	f := common.GetFile()
	defer f.Close()

	p1, p2, err := parse(f)
	if err != nil {
		log.Fatal(err)
	}

	res1, dur1 := common.TimeIt(func() int {
		return part1(p1, p2)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)
}

func part1(p1Start, p2Start int) int {
	dice := Dice{NextValue: 0}
	p1 := Player{Pos: p1Start}
	p2 := Player{Pos: p2Start}

	p1Turn := true

	for !p1.Won() && !p2.Won() {
		roll := dice.Roll3()
		if p1Turn {
			p1.Move(roll)
		} else {
			p2.Move(roll)
		}
		p1Turn = !p1Turn
	}

	if p1.Won() {
		return p2.Score * dice.Rolled
	}

	return p1.Score * dice.Rolled
}

func parse(r io.Reader) (int, int, error) {
	br := bufio.NewReader(r)

	var p1, p2 int

	if _, err := fmt.Fscanf(br, "Player 1 starting position: %d\n", &p1); err != nil {
		return 0, 0, err
	}
	if _, err := fmt.Fscanf(br, "Player 2 starting position: %d\n", &p2); err != nil {
		return 0, 0, err
	}

	return p1, p2, nil
}

func (p *Player) Move(roll int) {
	if (p.Pos+roll)%10 == 0 {
		p.Pos = 10
	} else {
		p.Pos = (p.Pos + roll) % 10
	}
	p.Score += p.Pos
}

func (p Player) Won() bool {
	return p.Score >= 1000
}

func (d *Dice) Roll3() int {
	d.Rolled += 3
	n1 := d.NextValue % 100
	n2 := (d.NextValue + 1) % 100
	n3 := (d.NextValue + 2) % 100
	d.NextValue = (d.NextValue + 3) % 100
	return n1 + 1 + n2 + 1 + n3 + 1
}
