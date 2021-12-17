package jurassicjigsaw

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Tile struct {
	ID      int
	Matrix  [][]string
	Aligned bool
}

type Point struct {
	X int
	Y int
}

type EdgeSet struct {
	Top    string
	Bottom string
	Left   string
	Right  string
}

func readLines(path string) ([]Tile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []Tile = make([]Tile, 0)

	r, err := regexp.Compile(`^Tile (\d+):$`)
	if err != nil {
		return nil, errors.New("couldn't parse regex expression")
	}

	scanner := bufio.NewScanner(file)
	var curIndex int
	var curMatrix [][]string = make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			result = append(result, Tile{ID: curIndex, Matrix: curMatrix})
			curMatrix = make([][]string, 0)
		} else if r.MatchString(line) {
			content := r.FindStringSubmatch(line)
			curIndex, err = strconv.Atoi(content[1])
			if err != nil {
				return nil, errors.New("error fetching tile ID")
			}
		} else {
			curMatrix = append(curMatrix, strings.Split(line, ""))
		}
	}
	result = append(result, Tile{ID: curIndex, Matrix: curMatrix})

	return result, nil
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func alreadyUsed(matrix map[Point]Tile, id int) bool {
	for _, v := range matrix {
		if v.ID == id {
			return true
		}
	}
	return false
}

func getEdgeSet(tile Tile) EdgeSet {
	var result EdgeSet
	var leftSlice, rightSlice []string = make([]string, 0), make([]string, 0)

	for i, _ := range tile.Matrix {
		leftSlice = append(leftSlice, tile.Matrix[i][0])
		rightSlice = append(rightSlice, tile.Matrix[i][len(tile.Matrix[i])-1])
	}

	result.Top = strings.Join(tile.Matrix[0], "")
	result.Bottom = strings.Join(tile.Matrix[len(tile.Matrix)-1], "")
	result.Left = strings.Join(leftSlice, "")
	result.Right = strings.Join(rightSlice, "")
	return result
}

func edgeList(tile Tile) []string {
	var result []string = make([]string, 0)
	var leftSlice, rightSlice []string = make([]string, 0), make([]string, 0)

	for i, _ := range tile.Matrix {
		leftSlice = append(leftSlice, tile.Matrix[i][0])
		rightSlice = append(rightSlice, tile.Matrix[i][len(tile.Matrix[i])-1])
	}

	result = append(result, strings.Join(tile.Matrix[0], ""))
	result = append(result, strings.Join(tile.Matrix[len(tile.Matrix)-1], ""))
	result = append(result, strings.Join(leftSlice, ""))
	result = append(result, strings.Join(rightSlice, ""))
	return result
}

func flipMatrix(data [][]string) [][]string {
	var result [][]string = make([][]string, len(data))
	for i, _ := range data {
		result[i] = make([]string, len(data[i]))
		for j := 0; j < len(data[i]); j++ {
			result[i][len(data[i])-j-1] = data[i][j]
		}
	}
	return result
}

func rotateMatrix(data [][]string) [][]string {
	var result [][]string = make([][]string, len(data))
	for i, _ := range data {
		result[i] = make([]string, len(data[i]))
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			result[j][len(data)-i-1] = data[i][j]
		}
	}
	return result
}

func flipTile(tile Tile) Tile {
	tile.Matrix = flipMatrix(tile.Matrix)
	return tile
}

func rotateTile(tile Tile) Tile {
	tile.Matrix = rotateMatrix(tile.Matrix)
	return tile
}

func findById(data []Tile, id int) Tile {
	for _, cur := range data {
		if cur.ID == id {
			return cur
		}
	}
	return Tile{}
}

func connections(data []Tile) map[int][]int {
	var result map[int][]int = make(map[int][]int)
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			iEdges := edgeList(data[i])
			jEdges := edgeList(data[j])
			for _, iEdge := range iEdges {
				for _, jEdge := range jEdges {
					if iEdge == jEdge || iEdge == reverse(jEdge) {
						result[data[i].ID] = append(result[data[i].ID], data[j].ID)
						result[data[j].ID] = append(result[data[j].ID], data[i].ID)
					}
				}
			}
		}
	}
	return result
}

func contains(data []int, value int) bool {
	for _, v := range data {
		if v == value {
			return true
		}
	}
	return false
}

func build(matrix map[Point]Tile, x int, y int, tile Tile, pairs map[int][]int, data []Tile, side int) bool {
	if x < 0 || y < 0 || x >= side || y >= side {
		return false
	}
	for _, v := range matrix {
		if v.ID == tile.ID {
			return false
		}
	}
	if x > 0 {
		if !contains(pairs[tile.ID], matrix[Point{X: x - 1, Y: y}].ID) {
			return false
		}
	}
	if y > 0 {
		if !contains(pairs[tile.ID], matrix[Point{X: x, Y: y - 1}].ID) {
			return false
		}
	}
	matrix[Point{X: x, Y: y}] = tile
	if x == side-1 && y == side-1 {
		return true
	}
	nextX, nextY := 0, 0
	var pairsToUse []int
	if x == side-1 {
		nextX, nextY = 0, y+1
		pairsToUse = pairs[matrix[Point{X: 0, Y: y}].ID]
	} else {
		nextX, nextY = x+1, y
		pairsToUse = pairs[tile.ID]
	}
	for _, v := range pairsToUse {
		if build(matrix, nextX, nextY, findById(data, v), pairs, data, side) {
			return true
		}
	}
	delete(matrix, Point{X: x, Y: y})
	return false
}

func variations(tile Tile) []Tile {
	return []Tile{
		tile,
		flipTile(tile),
		rotateTile(tile),
		rotateTile(rotateTile(tile)),
		rotateTile(rotateTile(rotateTile(tile))),
		rotateTile(flipTile(tile)),
		rotateTile(rotateTile(flipTile(tile))),
		rotateTile(rotateTile(rotateTile(flipTile(tile)))),
	}
}

func align(matrix map[Point]Tile, side int) {
	topLeft, topLeftRight, topLeftBottom := matrix[Point{X: 0, Y: 0}], matrix[Point{X: 1, Y: 0}], matrix[Point{X: 0, Y: 1}]
	topLeftRightEdges, topLeftBottomEdges := getEdgeSet(topLeftRight), getEdgeSet(topLeftBottom)
	topLeftVariations := variations(topLeft)
	for _, cur := range topLeftVariations {
		rightAligned, bottomAligned := false, false
		curEdges := getEdgeSet(cur)
		if curEdges.Right == topLeftRightEdges.Left || curEdges.Right == topLeftRightEdges.Top || curEdges.Right == topLeftRightEdges.Right || curEdges.Right == topLeftRightEdges.Bottom || curEdges.Right == reverse(topLeftRightEdges.Left) || curEdges.Right == reverse(topLeftRightEdges.Top) || curEdges.Right == reverse(topLeftRightEdges.Right) || curEdges.Right == reverse(topLeftRightEdges.Bottom) {
			rightAligned = true
		}
		if curEdges.Bottom == topLeftBottomEdges.Left || curEdges.Bottom == topLeftBottomEdges.Top || curEdges.Bottom == topLeftBottomEdges.Right || curEdges.Bottom == topLeftBottomEdges.Bottom || curEdges.Bottom == reverse(topLeftBottomEdges.Left) || curEdges.Bottom == reverse(topLeftBottomEdges.Top) || curEdges.Bottom == reverse(topLeftBottomEdges.Right) || curEdges.Bottom == reverse(topLeftBottomEdges.Bottom) {
			bottomAligned = true
		}
		if rightAligned && bottomAligned {
			cur.Aligned = true
			matrix[Point{X: 0, Y: 0}] = cur
			break
		}
	}

	for y := 0; y < side; y++ {
		for x := 0; x < side; x++ {
			if x == 0 && y == 0 {
				continue
			}
			point := Point{X: x, Y: y}
			top, left := false, false
			for !top || !left {
				curTile := matrix[point]
				curEdges := getEdgeSet(curTile)
				if y == 0 {
					top = true
				} else {
					topEdges := getEdgeSet(matrix[Point{X: x, Y: y - 1}])
					if curEdges.Top == topEdges.Bottom {
						top = true
					} else if reverse(curEdges.Left) == topEdges.Bottom {
						matrix[point] = rotateTile(curTile)
						top = true
					} else if reverse(curEdges.Bottom) == topEdges.Bottom {
						matrix[point] = rotateTile(rotateTile(curTile))
						top = true
					} else if curEdges.Right == topEdges.Bottom {
						matrix[point] = rotateTile(rotateTile(rotateTile(curTile)))
						top = true
					} else if reverse(curEdges.Top) == topEdges.Bottom {
						matrix[point] = flipTile(curTile)
						top = true
					} else if reverse(curEdges.Right) == topEdges.Bottom {
						matrix[point] = rotateTile(flipTile(curTile))
						top = true
					} else if curEdges.Bottom == topEdges.Bottom {
						matrix[point] = rotateTile(rotateTile(flipTile(curTile)))
						top = true
					} else if curEdges.Left == topEdges.Bottom {
						matrix[point] = rotateTile(rotateTile(rotateTile(flipTile(curTile))))
						top = true
					} else {
						break
					}
				}
				if x == 0 {
					left = true
				} else {
					leftEdges := getEdgeSet(matrix[Point{X: x - 1, Y: y}])
					if curEdges.Left == leftEdges.Right {
						left = true
					} else if curEdges.Bottom == leftEdges.Right {
						matrix[point] = rotateTile(curTile)
						left = true
					} else if reverse(curEdges.Right) == leftEdges.Right {
						matrix[point] = rotateTile(rotateTile(curTile))
						left = true
					} else if reverse(curEdges.Top) == leftEdges.Right {
						matrix[point] = rotateTile(rotateTile(rotateTile(curTile)))
						left = true
					} else if curEdges.Right == leftEdges.Right {
						matrix[point] = flipTile(curTile)
						left = true
					} else if reverse(curEdges.Bottom) == leftEdges.Right {
						matrix[point] = rotateTile(flipTile(curTile))
						left = true
					} else if reverse(curEdges.Left) == leftEdges.Right {
						matrix[point] = rotateTile(rotateTile(flipTile(curTile)))
						left = true
					} else if curEdges.Top == leftEdges.Right {
						matrix[point] = rotateTile(rotateTile(rotateTile(flipTile(curTile))))
						left = true
					} else {
						break
					}
				}
				if top && left {
					curTile = matrix[point]
					curTile.Aligned = true
					matrix[point] = curTile
				}
			}
		}
	}
}

func removeEdges(tile Tile) Tile {
	var newMatrix [][]string = make([][]string, len(tile.Matrix)-2)
	for i, _ := range newMatrix {
		newMatrix[i] = make([]string, len(tile.Matrix[0])-2)
	}

	for i, a := 1, 0; i < len(tile.Matrix)-1; i, a = i+1, a+1 {
		for j, b := 1, 0; j < len(tile.Matrix[i])-1; j, b = j+1, b+1 {
			newMatrix[a][b] = tile.Matrix[i][j]
		}
	}

	tile.Matrix = newMatrix
	return tile
}

func joinTiles(matrix map[Point]Tile, size int) [][]string {
	var result [][]string = make([][]string, size*8)
	for i := 0; i < size*8; i++ {
		result[i] = make([]string, 0)
	}
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			tile := matrix[Point{X: x, Y: y}]
			for i := 0; i < 8; i++ {
				for _, cur := range tile.Matrix[i] {
					result[8*y+i] = append(result[8*y+i], cur)
				}
			}
		}
	}
	return result
}

func countHashtags(data [][]string) int {
	count := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == "#" {
				count++
			}
		}
	}
	return count
}

func mergeMatrix(data [][]string) string {
	result := ""
	for i, _ := range data {
		result = result + strings.Join(data[i], "")
	}
	return result
}

func lookForMonsters(data [][]string) (int, error) {
	hashTagsCount := countHashtags(data)

	rTop, err := regexp.Compile(`^.{18}\#`)
	if err != nil {
		return 0, errors.New("couldn't parse regex expression")
	}
	rMid, err := regexp.Compile(`\#.{4}\#\#.{4}\#\#.{4}\#\#\#`)
	if err != nil {
		return 0, errors.New("couldn't parse regex expression")
	}
	rBottom, err := regexp.Compile(`^.\#..\#..\#..\#..\#..\#`)
	if err != nil {
		return 0, errors.New("couldn't parse regex expression")
	}

	options := [][][]string{data, rotateMatrix(data), rotateMatrix(rotateMatrix(data)), rotateMatrix(rotateMatrix(rotateMatrix(data))), flipMatrix(data), rotateMatrix(flipMatrix(data)), rotateMatrix(rotateMatrix(flipMatrix(data))), rotateMatrix(rotateMatrix(rotateMatrix(flipMatrix(data))))}
	for _, cur := range options {
		monsters := 0
		for i := 1; i < len(cur)-1; i++ {
			line := strings.Join(cur[i], "")
			indexes := rMid.FindAllStringSubmatchIndex(line, -1)
			if indexes != nil {
				for _, idx := range indexes {
					topLine := strings.Join(cur[i-1], "")[idx[0]:]
					if rTop.MatchString(topLine) {
						bottomLine := strings.Join(cur[i+1], "")[idx[0]:]
						if rBottom.MatchString(bottomLine) {
							monsters++
						}
					}
				}
			}
		}
		if monsters > 0 {
			return hashTagsCount - (monsters * 15), nil
		}
	}

	return 0, errors.New("something went wrong")
}

func Challenge1(data []Tile) (int64, error) {
	pairs := connections(data)
	result := int64(1)

	for _, v := range data {
		if len(pairs[v.ID]) == 2 {
			result = result * int64(v.ID)
		}
	}

	return result, nil
}

func Challenge2(data []Tile) (int64, error) {
	pairs := connections(data)

	var corners []Tile = make([]Tile, 0)

	for k, v := range data {
		if len(pairs[v.ID]) == 2 {
			corners = append(corners, data[k])
		}
	}

	var matrix map[Point]Tile = make(map[Point]Tile)
	size := int(math.Sqrt(float64(len(data))))

	for _, cur := range corners {
		if build(matrix, 0, 0, cur, pairs, data, size) {
			break
		}
	}

	align(matrix, size)

	var newMatrix map[Point]Tile = make(map[Point]Tile)
	for k, v := range matrix {
		newMatrix[k] = removeEdges(v)
	}

	joined := joinTiles(newMatrix, size)
	result, err := lookForMonsters(joined)
	if err == nil {
		return int64(result), nil
	} else {
		return 0, err
	}
}

func Run() {
	fmt.Println("Day 20 - Jurassic Jigsaw")
	path := "jurassicjigsaw/input.txt"
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
