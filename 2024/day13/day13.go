package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Pos struct {
	X int
	Y int
}

type Game struct {
	ButtonA Pos
	ButtonB Pos
	Prize   Pos
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	games := parse(string(data))
	res1 := solve(games)
	fmt.Printf("Part1: %d\n", res1)

	for i := 0; i < len(games); i++ {
		games[i].Prize.X += 10_000_000_000_000
		games[i].Prize.Y += 10_000_000_000_000
	}
	res2 := solve(games)
	fmt.Printf("Part2: %d\n", res2)
}

func solve(games []Game) int {
	var res int
	for _, game := range games {
		X1 := float64(game.ButtonA.X)
		Y1 := float64(game.ButtonA.Y)
		X2 := float64(game.ButtonB.X)
		Y2 := float64(game.ButtonB.Y)
		pX := float64(game.Prize.X)
		pY := float64(game.Prize.Y)

		xRatio := Y1 / X1

		bFloat := (pY - xRatio*pX) / (-xRatio*X2 + Y2)
		aFloat := (pX - X2*bFloat) / X1
		a := int(math.Round(aFloat))
		b := int(math.Round(bFloat))
		xSol := a*game.ButtonA.X + b*game.ButtonB.X
		ySol := a*game.ButtonA.Y + b*game.ButtonB.Y

		if xSol == game.Prize.X && ySol == game.Prize.Y {
			res += a*3 + b
		}
	}
	return res
}

func parse(content string) []Game {
	var games []Game

	for _, gameStr := range strings.Split(content, "\n\n") {
		games = append(games, parseGame(gameStr))
	}

	return games
}

func parseGame(str string) Game {
	lines := strings.Split(str, "\n")
	return Game{
		ButtonA: intoPos(numberRegex.FindAllString(lines[0], 2)),
		ButtonB: intoPos(numberRegex.FindAllString(lines[1], 2)),
		Prize:   intoPos(numberRegex.FindAllString(lines[2], 2)),
	}
}

var numberRegex = regexp.MustCompile(`\d+`)

func intoPos(numStr []string) Pos {
	x, err := strconv.Atoi(numStr[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(numStr[1])
	if err != nil {
		log.Fatal(err)
	}
	return Pos{X: x, Y: y}
}
