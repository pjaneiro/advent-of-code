package reactorreboot

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Cuboid struct {
	Type bool
	XMin int64
	XMax int64
	YMin int64
	YMax int64
	ZMin int64
	ZMax int64
}

type Point struct {
	X int64
	Y int64
	Z int64
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func (c Cuboid) volume() int64 {
	return (c.XMax - c.XMin + 1) * (c.YMax - c.YMin + 1) * (c.ZMax - c.ZMin + 1)
}

func (c Cuboid) intersection(other Cuboid, typeval bool) Cuboid {
	return Cuboid{
		Type: typeval,
		XMin: max(c.XMin, other.XMin),
		XMax: min(c.XMax, other.XMax),
		YMin: max(c.YMin, other.YMin),
		YMax: min(c.YMax, other.YMax),
		ZMin: max(c.ZMin, other.ZMin),
		ZMax: min(c.ZMax, other.ZMax),
	}
}

func (c Cuboid) valid() bool {
	return c.XMin <= c.XMax && c.YMin <= c.YMax && c.ZMin <= c.ZMax
}

func readLines(path string) ([]Cuboid, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r, err := regexp.Compile(`^(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)$`)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	var result []Cuboid = make([]Cuboid, 0)
	for scanner.Scan() {
		var elem Cuboid
		data := r.FindStringSubmatch(scanner.Text())
		if data[1] == "on" {
			elem.Type = true
		} else {
			elem.Type = false
		}
		elem.XMin, _ = strconv.ParseInt(data[2], 10, 64)
		elem.XMax, _ = strconv.ParseInt(data[3], 10, 64)
		elem.YMin, _ = strconv.ParseInt(data[4], 10, 64)
		elem.YMax, _ = strconv.ParseInt(data[5], 10, 64)
		elem.ZMin, _ = strconv.ParseInt(data[6], 10, 64)
		elem.ZMax, _ = strconv.ParseInt(data[7], 10, 64)

		result = append(result, elem)
	}

	return result, nil
}

func Challenge1(data []Cuboid) (int64, error) {
	var mapping map[Point]struct{} = make(map[Point]struct{})

	for _, command := range data {
		if command.XMin > 50 || command.XMax < -50 || command.YMin > 50 || command.YMax < -50 || command.ZMin > 50 || command.ZMax < -50 {
			continue
		}
		for x := command.XMin; x <= command.XMax; x++ {
			if x < -50 || x > 50 {
				continue
			}
			for y := command.YMin; y <= command.YMax; y++ {
				if y < -50 || y > 50 {
					continue
				}
				for z := command.ZMin; z <= command.ZMax; z++ {
					if z < -50 || z > 50 {
						continue
					}

					if command.Type {
						mapping[Point{X: x, Y: y, Z: z}] = struct{}{}
					} else {
						if _, ok := mapping[Point{X: x, Y: y, Z: z}]; ok {
							delete(mapping, Point{X: x, Y: y, Z: z})
						}
					}
				}
			}
		}
	}

	return int64(len(mapping)), nil
}

func Challenge2(data []Cuboid) (int64, error) {
	var resultCubes []Cuboid = make([]Cuboid, 0)
	for _, cur := range data {
		var toAdd []Cuboid = make([]Cuboid, 0)
		if cur.Type {
			toAdd = append(toAdd, cur)
		}
		for _, other := range resultCubes {
			if intersect := cur.intersection(other, !other.Type); intersect.valid() {
				toAdd = append(toAdd, intersect)
			}
		}
		resultCubes = append(resultCubes, toAdd...)
	}
	var result int64 = 0
	for _, cur := range resultCubes {
		if cur.Type {
			result = result + cur.volume()
		} else {
			result = result - cur.volume()
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 22 - Reactor Reboot")
	path := "reactorreboot/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int64
	result, err = Challenge1(data)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
