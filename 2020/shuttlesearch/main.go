package shuttlesearch

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) (int, []string, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	goal, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, nil, err
	}
	scanner.Scan()
	line := scanner.Text()
	ids := strings.Split(line, ",")

	return goal, ids, nil
}

func Challenge1(goal int, ids []string) (int, error) {
	diff := math.MaxInt32
	best := 0

	for _, v := range ids {
		id, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		i := id
		for i < goal {
			i += id
		}
		// i := id + (goal - id%goal)
		tmpdiff := i - goal
		if tmpdiff < diff {
			diff = tmpdiff
			best = id
		}
	}

	return best * diff, nil
}

func Challenge2(data []string) (int64, error) {
	var ids []int64
	for _, v := range data {
		id, _ := strconv.Atoi(v)
		ids = append(ids, int64(id))
	}

	result, step := ids[0], ids[0]
	for i := 1; i < len(ids); i++ {
		if ids[i] == 0 {
			continue
		}
		n := ids[i] - (int64(i) % ids[i])

		for result%ids[i] != n {
			result += step
		}
		step *= ids[i]
	}

	return result, nil
}

// Version with sorted data
/*func Challenge2(data []string) (int64, error) {
	var ids []int64
	for _, v := range data {
		id, _ := strconv.Atoi(v)
		ids = append(ids, int64(id))
	}

	var tmp map[int]int = make(map[int]int, len(ids))
	for i := range ids {
		if ids[i] != 0 {
			tmp[int(ids[i])] = i
		}
	}
	tmp2 := make([]int, 0)
	ids = make([]int64, 0)
	is := make([]int64, 0)

	for k, _ := range tmp {
		tmp2 = append(tmp2, k)
	}
	sort.Ints(tmp2)
	for i, j := 0, len(tmp2)-1; i < j; i, j = i+1, j-1 {
		tmp2[i], tmp2[j] = tmp2[j], tmp2[i]
	}
	for _, k := range tmp2 {
		ids = append(ids, int64(k))
		is = append(is, int64(tmp[k]))
	}

	result := ids[0] - is[0]
	step := ids[0]
	for i := 1; i < len(ids); i++ {
		var n int64
		if is[i] == 0 {
			n = 0
		} else {
			n = ids[i] - is[i]
			for n < 0 {
				n += ids[i]
			}
		}

		for result%ids[i] != n {
			result += step
		}
		step *= ids[i]
	}

	return result, nil
}*/

func Run() {
	fmt.Println("Day 13 - Shuttle Search")
	path := "shuttlesearch/input.txt"
	goal, ids, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(goal, ids)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	var result2 int64
	result2, err = Challenge2(ids)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result2)
	}
}
