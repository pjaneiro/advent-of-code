package lobbylayout

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

type Coord struct {
	x int
	y int
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string = make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func parseTasks(data []string) []map[string]int {
	r, err := regexp.Compile(`(se|sw|ne|nw|(?:[^nsew]?e)|(?:[^nsew]?w))`)
	if err != nil {
		return nil
	}
	var result []map[string]int = make([]map[string]int, 0)
	for _, cur := range data {
		var curMap map[string]int = make(map[string]int)
		found := r.FindAllString(cur, -1)
		for _, match := range found {
			curMap[match]++
		}
		result = append(result, curMap)
	}
	return result
}

func even(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func Challenge1(data []string) (int, error) {
	parsedTasks := parseTasks(data)

	var flips map[Coord]struct{} = make(map[Coord]struct{})

	for _, curTask := range parsedTasks {
		e, ne, se, w, nw, sw := curTask["e"], curTask["ne"], curTask["se"], curTask["w"], curTask["nw"], curTask["sw"]
		point := Coord{x: (2 * e) - (2 * w) + ne + se - nw - sw, y: ne + nw - se - sw}

		if _, ok := flips[point]; ok {
			delete(flips, point)
		} else {
			flips[point] = struct{}{}
		}
	}

	return len(flips), nil
}

func Challenge2(data []string) (int, error) {
	parsedTasks := parseTasks(data)

	var flips map[Coord]struct{} = make(map[Coord]struct{})

	minX, maxX, minY, maxY := math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32

	for _, curTask := range parsedTasks {
		e, ne, se, w, nw, sw := curTask["e"], curTask["ne"], curTask["se"], curTask["w"], curTask["nw"], curTask["sw"]
		x, y := (2*e)-(2*w)+ne+se-nw-sw, ne+nw-se-sw

		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}

		point := Coord{x: x, y: y}
		if _, ok := flips[point]; ok {
			delete(flips, point)
		} else {
			flips[point] = struct{}{}
		}
	}

	for round := 0; round < 100; round++ {
		var newFlips map[Coord]struct{} = make(map[Coord]struct{})
		newMinX, newMaxX, newMinY, newMaxY := minX, maxX, minY, maxY

		for i := minX - 2; i <= maxX+1; i++ {
			for j := minY - 2; j <= maxY+1; j++ {
				if even(i) != even(j) {
					continue
				}

				var color rune

				if _, ok := flips[Coord{x: i, y: j}]; ok {
					color = 'b'
				} else {
					color = 'w'
				}

				blacks := 0

				neighbours := []Coord{
					Coord{x: i - 2, y: j},     // w
					Coord{x: i + 2, y: j},     // e
					Coord{x: i + 1, y: j + 1}, // ne
					Coord{x: i + 1, y: j - 1}, // se
					Coord{x: i - 1, y: j + 1}, // nw
					Coord{x: i - 1, y: j - 1}, // sw
				}

				for _, curN := range neighbours {
					if _, ok := flips[curN]; ok {
						blacks++
					}
					if curN.x < newMinX {
						newMinX = curN.x
					}
					if curN.x > newMaxX {
						newMaxX = curN.x
					}
					if curN.y < newMinY {
						newMinY = curN.y
					}
					if curN.y > newMaxY {
						newMaxY = curN.y
					}
				}

				if (color == 'b' && (blacks == 1 || blacks == 2)) || (color == 'w' && blacks == 2) {
					newFlips[Coord{x: i, y: j}] = struct{}{}
				}
			}
		}
		flips = newFlips
		minX, maxX, minY, maxY = newMinX, newMaxX, newMinY, newMaxY
	}

	return len(flips), nil
}

func Run() {
	fmt.Println("Day 24 - Lobby Layout")
	path := "lobbylayout/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
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
