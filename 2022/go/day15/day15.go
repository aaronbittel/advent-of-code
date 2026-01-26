package main

import (
	"AOC2022/internal/common"
	"bufio"
	"fmt"
	"io"
	"log"
	"slices"
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
		return part1(sensors, 10)
	})
	fmt.Printf("Part1: %d, took %s\n", res1, dur1)

	res2, dur2 := common.TimeIt(func() int {
		return part2(sensors, 4_000_000)
	})
	fmt.Printf("Part2: %d, took %s\n", res2, dur2)
}

type Interval struct {
	Start, End int
}

func part2(sensors []Sensor, maxValue int) int {
	var resY, resX int

outer:
	for row := range maxValue {
		intervals := make([]Interval, 0, len(sensors))

		for _, sensor := range sensors {
			distToRow := common.AbsInt(sensor.Y - row)
			totalDist := sensor.Dist()
			dist := totalDist - distToRow
			if dist < 0 {
				continue
			}
			startX := max(sensor.X-dist, 0)
			endX := min(sensor.X+dist, maxValue)
			intervals = append(intervals, Interval{Start: startX, End: endX})
		}
		slices.SortFunc(intervals, func(a, b Interval) int {
			if a.Start < b.Start {
				return -1
			}
			if a.Start > b.Start {
				return 1
			}
			return 0
		})

		x, hasGap := FindGap(intervals)
		if hasGap {
			resY = row
			resX = x
			break outer
		}
	}

	return resX*4_000_000 + resY
}

func FindGap(intervals []Interval) (int, bool) {
	res := intervals[0]

	for _, interval := range intervals[1:] {
		if res.End < interval.Start {
			return res.End + 1, true
		}
		if res.End < interval.End {
			res.End = interval.End
		}
	}
	return 0, false
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
