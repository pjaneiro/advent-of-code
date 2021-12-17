package chiton

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Item struct {
	coords Point
	cost   int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	result := (*pq)[0]
	*pq = (*pq)[1:]
	return result
}

func readLines(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]int = make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), "")
		parsedLine := make([]int, len(splitLine))
		for k, v := range splitLine {
			num, err := strconv.Atoi(v)
			if err != nil {
				return result, err
			}
			parsedLine[k] = num
		}
		result = append(result, parsedLine)
	}

	return result, nil
}

func findRisk(data [][]int, times int) int {
	var result int
	origSize := len(data)
	fullSize := origSize * times

	var queue PriorityQueue
	heap.Init(&queue)

	var distances map[Point]int = make(map[Point]int, fullSize*fullSize)
	var previous map[Point]Point = make(map[Point]Point, fullSize*fullSize)

	start := Point{X: 0, Y: 0}
	goal := Point{X: fullSize - 1, Y: fullSize - 1}

	for i := 0; i < fullSize; i++ {
		for j := 0; j < fullSize; j++ {
			distances[Point{X: j, Y: i}] = math.MaxInt32
		}
	}

	distances[start] = 0
	heap.Push(&queue, &Item{coords: start, cost: 0})

	for queue.Len() != 0 {
		current := (*(heap.Pop(&queue).(*Item)))
		x, y := current.coords.X, current.coords.Y

		neighbours := []Point{Point{X: x, Y: y - 1}, Point{X: x, Y: y + 1}, Point{X: x - 1, Y: y}, Point{X: x + 1, Y: y}}
		for _, n := range neighbours {
			if n.X < 0 || n.Y < 0 || n.X >= fullSize || n.Y >= fullSize {
				continue
			}

			tmp := distances[current.coords] + ((data[n.Y%origSize][n.X%origSize]+(n.X/origSize)+(n.Y/origSize)-1)%9 + 1)
			if tmp < distances[n] {
				distances[n] = tmp
				previous[n] = current.coords
				heap.Push(&queue, &Item{coords: n, cost: tmp})
			}
		}
	}

	for cur := goal; cur != start; cur = previous[cur] {
		result += (data[cur.Y%origSize][cur.X%origSize]+(cur.X/origSize)+(cur.Y/origSize)-1)%9 + 1
	}
	return result
}

func Challenge1(data [][]int) (int, error) {
	return findRisk(data, 1), nil
}

func Challenge2(data [][]int) (int, error) {
	return findRisk(data, 5), nil
}

func Run() {
	fmt.Println("Day 15 - Chiton")
	path := "chiton/input.txt"
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
