package beaconscanner

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Beacon struct {
	Pos   [3]int
	Dists []int
}

type Scanner struct {
	ID      int
	Beacons []Beacon
	Pos     [3]int
	Rot     [3][3]int
	Located bool
}

func readLines(path string) ([]Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []Scanner = make([]Scanner, 0)
outerloop:
	for i := 0; i <= 36; i++ {
		var curScanner Scanner = Scanner{ID: i}
		scanner.Scan()
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				result = append(result, curScanner)
				continue outerloop
			}
			fields := strings.Split(line, ",")
			x, _ := strconv.Atoi(fields[0])
			y, _ := strconv.Atoi(fields[1])
			z, _ := strconv.Atoi(fields[2])
			curScanner.Beacons = append(curScanner.Beacons, Beacon{Pos: [3]int{x, y, z}})
		}
		result = append(result, curScanner)
	}

	return result, nil
}

func rotationMatrices() [24][3][3]int {
	return [24][3][3]int{
		[3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}, // identity
		[3][3]int{[3]int{1, 0, 0}, [3]int{0, 0, -1}, [3]int{0, 1, 0}},
		[3][3]int{[3]int{1, 0, 0}, [3]int{0, -1, 0}, [3]int{0, 0, -1}},
		[3][3]int{[3]int{1, 0, 0}, [3]int{0, 0, 1}, [3]int{0, -1, 0}},

		[3][3]int{[3]int{0, -1, 0}, [3]int{1, 0, 0}, [3]int{0, 0, 1}},
		[3][3]int{[3]int{0, 0, 1}, [3]int{1, 0, 0}, [3]int{0, 1, 0}},
		[3][3]int{[3]int{0, 1, 0}, [3]int{1, 0, 0}, [3]int{0, 0, -1}},
		[3][3]int{[3]int{0, 0, -1}, [3]int{1, 0, 0}, [3]int{0, -1, 0}},

		[3][3]int{[3]int{-1, 0, 0}, [3]int{0, -1, 0}, [3]int{0, 0, 1}},
		[3][3]int{[3]int{-1, 0, 0}, [3]int{0, 0, -1}, [3]int{0, -1, 0}},
		[3][3]int{[3]int{-1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, -1}},
		[3][3]int{[3]int{-1, 0, 0}, [3]int{0, 0, 1}, [3]int{0, 1, 0}},

		[3][3]int{[3]int{0, 1, 0}, [3]int{-1, 0, 0}, [3]int{0, 0, 1}},
		[3][3]int{[3]int{0, 0, 1}, [3]int{-1, 0, 0}, [3]int{0, -1, 0}},
		[3][3]int{[3]int{0, -1, 0}, [3]int{-1, 0, 0}, [3]int{0, 0, -1}},
		[3][3]int{[3]int{0, 0, -1}, [3]int{-1, 0, 0}, [3]int{0, 1, 0}},

		[3][3]int{[3]int{0, 0, -1}, [3]int{0, 1, 0}, [3]int{1, 0, 0}},
		[3][3]int{[3]int{0, 1, 0}, [3]int{0, 0, 1}, [3]int{1, 0, 0}},
		[3][3]int{[3]int{0, 0, 1}, [3]int{0, -1, 0}, [3]int{1, 0, 0}},
		[3][3]int{[3]int{0, -1, 0}, [3]int{0, 0, -1}, [3]int{1, 0, 0}},

		[3][3]int{[3]int{0, 0, -1}, [3]int{0, -1, 0}, [3]int{-1, 0, 0}},
		[3][3]int{[3]int{0, -1, 0}, [3]int{0, 0, 1}, [3]int{-1, 0, 0}},
		[3][3]int{[3]int{0, 0, 1}, [3]int{0, 1, 0}, [3]int{-1, 0, 0}},
		[3][3]int{[3]int{0, 1, 0}, [3]int{0, 0, -1}, [3]int{-1, 0, 0}},
	}
}

func findRotationIndex(matrix [3][3]int) int {
	for result, cur := range rotationMatrices() {
		matches := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if cur[i][j] == matrix[i][j] {
					matches++
				}
			}
		}
		if matches == 9 {
			return result
		}
	}
	return -1
}

func multiplyMatrixVector(matrix [3][3]int, vector [3]int) [3]int {
	return [3]int{
		matrix[0][0]*vector[0] + matrix[0][1]*vector[1] + matrix[0][2]*vector[2],
		matrix[1][0]*vector[0] + matrix[1][1]*vector[1] + matrix[1][2]*vector[2],
		matrix[2][0]*vector[0] + matrix[2][1]*vector[1] + matrix[2][2]*vector[2],
	}
}

func multiplyMatrixMatrix(matrix1 [3][3]int, matrix2 [3][3]int) [3][3]int {
	var result [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return result
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func translateVector(vector1 [3]int, vector2 [3]int) [3]int {
	return [3]int{vector1[0] + vector2[0], vector1[1] + vector2[1], vector1[2] + vector2[2]}
}

func distance(vector1 [3]int, vector2 [3]int) int {
	return abs(vector2[0]-vector1[0]) + abs(vector2[1]-vector1[1]) + abs(vector2[2]-vector1[2])
}

func addition(vector1 [3]int, vector2 [3]int) [3]int {
	return [3]int{vector1[0] + vector2[0], vector1[1] + vector2[1], vector1[2] + vector2[2]}
}

func difference(vector1 [3]int, vector2 [3]int) [3]int {
	return [3]int{vector1[0] - vector2[0], vector1[1] - vector2[1], vector1[2] - vector2[2]}
}

func equal(vector1 [3]int, vector2 [3]int) bool {
	return vector1[0] == vector2[0] && vector1[1] == vector2[1] && vector1[2] == vector2[2]
}

func contains(data []int, val int) bool {
	for _, cur := range data {
		if cur == val {
			return true
		}
	}
	return false
}

func intersection(data1 []int, data2 []int) []int {
	var result []int = make([]int, 0)
	for _, cur := range data1 {
		if contains(data2, cur) {
			result = append(result, cur)
		}
	}
	return result
}

func Challenge1(data []Scanner) (int, error) {
	var absPositions map[[3]int]struct{} = make(map[[3]int]struct{})
	for i, _ := range data {
		for j, _ := range data[i].Beacons {
			for k, _ := range data[i].Beacons {
				if k == j {
					continue
				}
				data[i].Beacons[j].Dists = append(data[i].Beacons[j].Dists, distance(data[i].Beacons[j].Pos, data[i].Beacons[k].Pos))
			}
		}
	}
	data[0].Pos = [3]int{0, 0, 0}
	data[0].Rot = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	data[0].Located = true
	positionedScanners := 0
	for i := 0; i < len(data); i++ {
		if data[i].Located {
			positionedScanners++
		}
	}
outerloop:
	for positionedScanners < len(data) {
		for leftIdx := 0; leftIdx < len(data); leftIdx++ {
			if data[leftIdx].Located {
				continue
			}

			for rightIdx := 0; rightIdx < len(data); rightIdx++ {
				if leftIdx == rightIdx {
					continue
				}
				if !data[rightIdx].Located {
					continue
				}

				var sel1, sel2 [][3]int = make([][3]int, 3), make([][3]int, 3)
				var picked int = 0

			innerloop:
				for i := 0; i < len(data[rightIdx].Beacons) && picked < 3; i++ {
					for j := 0; j < len(data[leftIdx].Beacons); j++ {
						if d := intersection(data[rightIdx].Beacons[i].Dists, data[leftIdx].Beacons[j].Dists); len(d) >= 11 {
							sel1[picked][0] = data[rightIdx].Beacons[i].Pos[0]
							sel1[picked][1] = data[rightIdx].Beacons[i].Pos[1]
							sel1[picked][2] = data[rightIdx].Beacons[i].Pos[2]
							sel2[picked][0] = data[leftIdx].Beacons[j].Pos[0]
							sel2[picked][1] = data[leftIdx].Beacons[j].Pos[1]
							sel2[picked][2] = data[leftIdx].Beacons[j].Pos[2]
							picked++
							continue innerloop
						}
					}
				}
				if picked < 3 {
					continue
				}
				vec_0_1_2, vec_0_1_3 := difference(sel1[1], sel1[0]), difference(sel1[2], sel1[0])
				for _, rot := range rotationMatrices() {
					tmp1, tmp2, tmp3 := multiplyMatrixVector(rot, sel2[0]), multiplyMatrixVector(rot, sel2[1]), multiplyMatrixVector(rot, sel2[2])
					vec_1_1_2, vec_1_1_3 := difference(tmp2, tmp1), difference(tmp3, tmp1)
					if equal(vec_0_1_2, vec_1_1_2) && equal(vec_0_1_3, vec_1_1_3) {
						data[leftIdx].Rot = multiplyMatrixMatrix(data[rightIdx].Rot, rot)
						data[leftIdx].Pos = addition(multiplyMatrixVector(data[rightIdx].Rot, difference(sel1[0], tmp1)), data[rightIdx].Pos)
						data[leftIdx].Located = true
						positionedScanners++
						continue outerloop
					}
				}
			}
		}
	}
	for idx := 0; idx < len(data); idx++ {
		for _, cur := range data[idx].Beacons {
			absPos := addition(multiplyMatrixVector(data[idx].Rot, cur.Pos), data[idx].Pos)
			absPositions[absPos] = struct{}{}
		}
	}
	return len(absPositions), nil
}

func Challenge2(data []Scanner) (int, error) {
	for i, _ := range data {
		for j, _ := range data[i].Beacons {
			for k, _ := range data[i].Beacons {
				if k == j {
					continue
				}
				data[i].Beacons[j].Dists = append(data[i].Beacons[j].Dists, distance(data[i].Beacons[j].Pos, data[i].Beacons[k].Pos))
			}
		}
	}
	data[0].Pos = [3]int{0, 0, 0}
	data[0].Rot = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	data[0].Located = true
	positionedScanners := 0
	for i := 0; i < len(data); i++ {
		if data[i].Located {
			positionedScanners++
		}
	}
outerloop:
	for positionedScanners < len(data) {
		for leftIdx := 0; leftIdx < len(data); leftIdx++ {
			if data[leftIdx].Located {
				continue
			}

			for rightIdx := 0; rightIdx < len(data); rightIdx++ {
				if leftIdx == rightIdx {
					continue
				}
				if !data[rightIdx].Located {
					continue
				}

				var sel1, sel2 [][3]int = make([][3]int, 3), make([][3]int, 3)
				var picked int = 0

			innerloop:
				for i := 0; i < len(data[rightIdx].Beacons) && picked < 3; i++ {
					for j := 0; j < len(data[leftIdx].Beacons); j++ {
						if d := intersection(data[rightIdx].Beacons[i].Dists, data[leftIdx].Beacons[j].Dists); len(d) >= 11 {
							sel1[picked][0] = data[rightIdx].Beacons[i].Pos[0]
							sel1[picked][1] = data[rightIdx].Beacons[i].Pos[1]
							sel1[picked][2] = data[rightIdx].Beacons[i].Pos[2]
							sel2[picked][0] = data[leftIdx].Beacons[j].Pos[0]
							sel2[picked][1] = data[leftIdx].Beacons[j].Pos[1]
							sel2[picked][2] = data[leftIdx].Beacons[j].Pos[2]
							picked++
							continue innerloop
						}
					}
				}
				if picked < 3 {
					continue
				}
				vec_0_1_2, vec_0_1_3 := difference(sel1[1], sel1[0]), difference(sel1[2], sel1[0])
				for _, rot := range rotationMatrices() {
					tmp1, tmp2, tmp3 := multiplyMatrixVector(rot, sel2[0]), multiplyMatrixVector(rot, sel2[1]), multiplyMatrixVector(rot, sel2[2])
					vec_1_1_2, vec_1_1_3 := difference(tmp2, tmp1), difference(tmp3, tmp1)
					if equal(vec_0_1_2, vec_1_1_2) && equal(vec_0_1_3, vec_1_1_3) {
						data[leftIdx].Rot = multiplyMatrixMatrix(data[rightIdx].Rot, rot)
						data[leftIdx].Pos = addition(multiplyMatrixVector(data[rightIdx].Rot, difference(sel1[0], tmp1)), data[rightIdx].Pos)
						data[leftIdx].Located = true
						positionedScanners++
						continue outerloop
					}
				}
			}
		}
	}
	best := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if d := distance(data[i].Pos, data[j].Pos); d > best {
				best = d
			}
		}
	}
	return best, nil
}

func Run() {
	fmt.Println("Day 19 - Beacon Scanner")
	path := "beaconscanner/input.txt"
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
