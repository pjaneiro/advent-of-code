package beaconexclusionzone

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Coord struct {
	X, Y int
}

type Reading struct {
	Sensor, Beacon Coord
	Distance       int
}

func abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}

func manDistance(a Coord, b Coord) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func ReadLines(path string) ([]Reading, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r, err := regexp.Compile(`^Sensor at x=(\-?\d+), y=(\-?\d+): closest beacon is at x=(\-?\d+), y=(\-?\d+)$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	var result []Reading
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := r.FindStringSubmatch(line)
		var reading Reading
		if reading.Sensor.X, err = strconv.Atoi(data[1]); err != nil {
			return nil, errors.New("error parsing input")
		}
		if reading.Sensor.Y, err = strconv.Atoi(data[2]); err != nil {
			return nil, errors.New("error parsing input")
		}
		if reading.Beacon.X, err = strconv.Atoi(data[3]); err != nil {
			return nil, errors.New("error parsing input")
		}
		if reading.Beacon.Y, err = strconv.Atoi(data[4]); err != nil {
			return nil, errors.New("error parsing input")
		}
		reading.Distance = abs(reading.Sensor.X-reading.Beacon.X) + abs(reading.Sensor.Y-reading.Beacon.Y)
		result = append(result, reading)
	}

	return result, nil
}

func Challenge1(data []Reading, row int) (int, error) {
	var beacons map[Coord]struct{} = make(map[Coord]struct{})
	var sensors map[Coord]struct{} = make(map[Coord]struct{})
	var intervals [][]int = make([][]int, 0)

	minX, maxX := math.MaxInt, math.MinInt

	for _, reading := range data {
		x, y, dist := reading.Sensor.X, reading.Sensor.Y, reading.Distance

		beacons[reading.Beacon] = struct{}{}
		sensors[reading.Sensor] = struct{}{}

		up, down, left, right := Coord{X: x, Y: y - dist}, Coord{X: x, Y: y + dist}, Coord{X: x - dist, Y: y}, Coord{X: x + dist, Y: y}
		if up.Y > row || down.Y < row {
			continue
		}
		yDiff := abs(row - y)
		leftX, rightX := left.X+yDiff, right.X-yDiff
		if leftX < minX {
			minX = leftX
		}
		if rightX > maxX {
			maxX = rightX
		}
		intervals = append(intervals, []int{leftX, rightX})
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	result := 0
	curX := minX
	for _, interval := range intervals {
		if interval[1] < curX {
			continue
		}
		if curX < interval[0] {
			curX = interval[0]
		}
		result += interval[1] - curX + 1
		curX = interval[1] + 1
	}
	for k, _ := range beacons {
		if k.Y == row {
			result--
		}
	}
	for k, _ := range sensors {
		if k.Y == row {
			result--
		}
	}
	return result, nil
}

func Challenge2(data []Reading, maxCoord int) (int, error) {
	sort.Slice(data, func(i, j int) bool { return data[i].Distance > data[j].Distance })

	for y := 0; y <= maxCoord; y++ {
		for x := 0; x <= maxCoord; x++ {
			var hit Reading
			found := true
			for _, reading := range data {
				if manDistance(Coord{X: x, Y: y}, reading.Sensor) <= reading.Distance {
					hit = reading
					found = false
					break
				}
			}
			if found {
				return x*4000000 + y, nil
			}

			yDiff := abs(y - hit.Sensor.Y)
			next := hit.Sensor.X + hit.Distance - yDiff
			if next > maxCoord {
				break
			}
			x = next
		}
	}
	return 0, errors.New("something went wrong")
}

func Run() {
	fmt.Println("Day 15 - Beacon Exclusion Zone")
	path := "beaconexclusionzone/input.txt"

	data, err := ReadLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(data, 2000000)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(data, 4000000)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
