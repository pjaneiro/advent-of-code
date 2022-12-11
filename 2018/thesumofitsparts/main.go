package thesumofitsparts

import (
	"bufio"
	"fmt"
	"os"
)

type Dependency struct {
	From, To rune
}

func readLines(path string) ([]Dependency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []Dependency
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		dep := Dependency{From: line[5], To: line[36]}
		result = append(result, dep)
	}
	return result, nil
}

func adjMatrixFromData(data []Dependency, alphaSize int) [][]int {
	var adjMatrix [][]int = make([][]int, alphaSize)
	for i := 0; i < alphaSize; i++ {
		adjMatrix[i] = make([]int, alphaSize)
	}

	for _, line := range data {
		adjMatrix[line.From-'A'][line.To-'A'] = 1
	}
	return adjMatrix
}

func Challenge1(data []Dependency, alphaSize int) (string, error) {
	adjMatrix := adjMatrixFromData(data, alphaSize)

	var placed map[int]struct{} = make(map[int]struct{})
	var result []rune = make([]rune, 0, alphaSize)

	for len(placed) < alphaSize {
	outerloop:
		for cur := 0; cur < alphaSize; cur++ {
			if _, ok := placed[cur]; ok {
				continue
			}
			for i := 0; i < alphaSize; i++ {
				if adjMatrix[i][cur] != 0 {
					continue outerloop
				}
			}
			result = append(result, rune(cur+'A'))
			placed[cur] = struct{}{}
			for j := 0; j < alphaSize; j++ {
				adjMatrix[cur][j] = 0
			}
			break
		}
	}

	return string(result), nil
}

func Challenge2(data []Dependency, alphaSize int, nWorkers int, offset int) (int, error) {
	adjMatrix := adjMatrixFromData(data, alphaSize)
	var dependencies []int = make([]int, alphaSize)
	for _, line := range data {
		dependencies[line.To-'A']++
	}

	var placed map[int]struct{} = make(map[int]struct{})
	var timeline [][]rune = make([][]rune, 1)
	timeline[0] = make([]rune, nWorkers)

	result := 0

	for second := 0; len(placed) < alphaSize; second++ {
		for worker := 0; second != 0 && worker < nWorkers; worker++ {
			if timeline[second][worker] == 0 && timeline[second-1][worker] != 0 {
				letter := int(timeline[second-1][worker] - 'A')
				placed[letter] = struct{}{}
				for i := 0; i < alphaSize; i++ {
					if adjMatrix[letter][i] != 0 {
						adjMatrix[letter][i] = 0
						dependencies[i]--
					}
				}
			}
		}
		for worker := 0; worker < nWorkers; worker++ {
			if timeline[second][worker] != 0 {
				continue
			}
			for i := 0; i < alphaSize; i++ {
				if _, ok := placed[i]; !ok && dependencies[i] == 0 {
					placed[i] = struct{}{}
					for j := second; j <= second+offset+i; j++ {
						if j >= len(timeline)-1 {
							timeline = append(timeline, make([]rune, nWorkers))
						}
						timeline[j][worker] = rune(i) + 'A'
						if j > result {
							result = j
						}
					}
					break
				}
			}
		}
	}

	return result + 1, nil
}

func Run() {
	fmt.Println("Day 6 - The Sum of Its Parts")
	path := "thesumofitsparts/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result1 string
	result1, err = Challenge1(data, 26)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %s\n", result1)
	}

	var result2 int
	result2, err = Challenge2(data, 26, 5, 60)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result2)
	}
}
