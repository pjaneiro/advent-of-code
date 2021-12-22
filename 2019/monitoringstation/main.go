package monitoringstation

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Asteroid struct {
	X         int
	Y         int
	Angle     float64
	Distance  int
	Destroyed bool
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data, nil
}

func abs(data int) int {
	if data < 0 {
		return 0 - data
	}
	return data
}

func gcd(val1 int, val2 int) int {
	for val2 != 0 {
		tmp := val2
		val2 = val1 % val2
		val1 = tmp
	}
	return val1
}

func Challenge1(data []string) (int, error) {
	best := 0
	for y1 := 0; y1 < len(data); y1++ {
		for x1 := 0; x1 < len(data[y1]); x1++ {
			if data[y1][x1] != '#' {
				continue
			}
			cur := 0

			for y2 := 0; y2 < len(data); y2++ {
				for x2 := 0; x2 < len(data[y2]); x2++ {
					if (x2 == x1 && y2 == y1) || data[y2][x2] != '#' {
						continue
					}

					dx, dy := x2-x1, y2-y1
					if dx == 0 {
						dy = dy / abs(dy)
					} else if dy == 0 {
						dx = dx / abs(dx)
					} else {
						d := gcd(abs(dx), abs(dy))
						dx, dy = dx/d, dy/d
					}

					blocked := false
					for newx, newy := x1+dx, y1+dy; newx != x2 || newy != y2; newx, newy = newx+dx, newy+dy {
						if data[newy][newx] == '#' {
							blocked = true
							break
						}
					}
					if !blocked {
						cur++
					}
				}
			}

			if cur > best {
				best = cur
			}
		}
	}
	return best, nil
}

func Challenge2(data []string) (int, error) {
	best := 0
	var seenAsteroids []Asteroid
	for y1 := 0; y1 < len(data); y1++ {
		for x1 := 0; x1 < len(data[y1]); x1++ {
			if data[y1][x1] != '#' {
				continue
			}
			cur := 0
			var curSeenAsteroids []Asteroid = make([]Asteroid, 0)

			for y2 := 0; y2 < len(data); y2++ {
				for x2 := 0; x2 < len(data[y2]); x2++ {
					if (x2 == x1 && y2 == y1) || data[y2][x2] != '#' {
						continue
					}

					dx, dy := x2-x1, y2-y1
					if dx == 0 {
						dy = dy / abs(dy)
					} else if dy == 0 {
						dx = dx / abs(dx)
					} else {
						d := gcd(abs(dx), abs(dy))
						dx, dy = dx/d, dy/d
					}

					blocked := false
					for newx, newy := x1+dx, y1+dy; newx != x2 || newy != y2; newx, newy = newx+dx, newy+dy {
						if data[newy][newx] == '#' {
							blocked = true
							break
						}
					}
					if !blocked {
						cur++
					}
					angle := 180 - ((180 / math.Pi) * math.Atan2(float64(x2-x1), float64(y2-y1)))
					curSeenAsteroids = append(curSeenAsteroids, Asteroid{X: x2, Y: y2, Angle: angle, Distance: abs(x2-x1) + abs(y2-y1), Destroyed: false})
				}
			}
			if cur > best {
				best = cur
				seenAsteroids = curSeenAsteroids
			}
		}
	}

	sort.Slice(seenAsteroids, func(a int, b int) bool {
		if seenAsteroids[a].Angle == seenAsteroids[b].Angle {
			return seenAsteroids[a].Distance < seenAsteroids[b].Distance
		}
		return seenAsteroids[a].Angle < seenAsteroids[b].Angle
	})

	var lastAngle float64 = seenAsteroids[len(seenAsteroids)-1].Angle
	var destroyed int = 0
	var asteroid Asteroid

	for i := 0; destroyed < 200; i = (i + 1) % len(seenAsteroids) {
		if seenAsteroids[i].Destroyed || seenAsteroids[i].Angle == lastAngle {
			continue
		}
		seenAsteroids[i].Destroyed = true
		lastAngle = seenAsteroids[i].Angle
		asteroid = seenAsteroids[i]
		destroyed++
	}

	return asteroid.X*100 + asteroid.Y, nil
}

func Run() {
	fmt.Println("Day 10 - Monitoring Station")
	path := "monitoringstation/input.txt"
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
