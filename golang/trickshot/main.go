package trickshot

import (
	// "errors"
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func getTrajectories(minX, maxX, minY, maxY int) (int, int, error) {
	bestY := math.MinInt32
	var workingSets map[Point]struct{} = (make(map[Point]struct{}))
	for velX := 0; velX <= maxX; velX++ {
		for velY := minY; velY < -minY; velY++ {
			curX, curY, curVelX, curVelY, curTopY := 0, 0, velX, velY, 0
			for true {
				curX, curY = curX+curVelX, curY+curVelY
				if curY > curTopY {
					curTopY = curY
				}
				if curVelX > 0 {
					curVelX--
				}
				curVelY--

				if curX >= minX && curX <= maxX && curY >= minY && curY <= maxY {
					workingSets[Point{X: velX, Y: velY}] = struct{}{}
					if curTopY > bestY {
						bestY = curTopY
					}
					break
				} else if curX > maxX || curY < minY || (curVelX <= 0 && curX < minX) {
					break
				}
			}
		}
	}

	return bestY, len(workingSets), nil
}

func Challenge1(minX, maxX, minY, maxY int) (int, error) {
	result, _, err := getTrajectories(minX, maxX, minY, maxY)

	return result, err
}

func Challenge2(minX, maxX, minY, maxY int) (int, error) {
	_, result, err := getTrajectories(minX, maxX, minY, maxY)

	return result, err
}

func Run() {
	fmt.Println("Day 17 - Trick Shot")

	var result int
	var err error
	result, err = Challenge1(248, 285, -85, -56)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(248, 285, -85, -56)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
