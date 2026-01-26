package main

import (
	"AOC2022/internal/common"
	"bufio"
	"fmt"
	"io"
	"log"
)

type Pos struct {
	Y, X int
}

type Sensor struct {
	Pos
	Beacon Pos
}

func main() {
	f := common.GetFile()
	defer f.Close()

	sensors := parse(f)

	res1, dur1 := common.TimeIt(func() int {
		return part1(sensors, 2_000_000)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)
}

func part1(sensors []Sensor, row int) int {
	res := make(map[int]struct{})
	beacons := make(map[int]struct{})

	for _, sensor := range sensors {
		if sensor.Beacon.Y == row {
			beacons[sensor.Beacon.X] = struct{}{}
		}
		distToRow := common.AbsInt(sensor.Y - row)
		totalDist := sensor.Dist()
		dist := totalDist - distToRow
		if dist < 0 {
			continue
		}
		for x := sensor.X - dist; x <= sensor.X+dist; x++ {
			if _, ok := beacons[x]; !ok {
				res[x] = struct{}{}
			}
		}
	}

	return len(res)
}

func (s Sensor) Dist() int {
	return common.AbsInt(s.Y-s.Beacon.Y) + common.AbsInt(s.X-s.Beacon.X)
}

func parse(r io.Reader) []Sensor {
	sensors := []Sensor{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var sY, sX, bY, bX int
		n, err := fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sX, &sY, &bX, &bY)
		if err != nil {
			log.Fatal(err)
		}
		if n != 4 {
			log.Fatal("illegal input")
		}
		sensors = append(sensors, Sensor{
			Pos:    Pos{Y: sY, X: sX},
			Beacon: Pos{Y: bY, X: bX},
		})
	}
	return sensors
}
