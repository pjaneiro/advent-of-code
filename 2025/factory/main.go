package factory

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	// Note: you'll need to download and install the Z3 C library
	z3 "github.com/aclements/go-z3/z3"
)

type machine struct {
	lights   []bool
	buttons  [][]int
	joltages []int
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func parseInput(data []string) ([]machine, error) {
	var result []machine = make([]machine, 0)

	rLights, err := regexp.Compile(`^\[([\.\#]+)\]`)
	if err != nil {
		return nil, err
	}
	rButtons, err := regexp.Compile(`\(([\,\d]+)\)`)
	if err != nil {
		return nil, err
	}
	rJoltage, err := regexp.Compile(`\{(.*)\}$`)
	if err != nil {
		return nil, err
	}

	for _, row := range data {
		var cur machine

		fLights := rLights.FindStringSubmatch(row)
		var lights []bool = make([]bool, len(fLights[1]))
		for l := range fLights[1] {
			if fLights[1][l] == '#' {
				lights[l] = true
			}
		}
		cur.lights = lights

		fButtons := rButtons.FindAllStringSubmatch(row, -1)
		var buttons [][]int = make([][]int, 0)
		for _, bSplit := range fButtons {
			var button []int = make([]int, 0)
			parts := strings.Split(bSplit[1], ",")
			for _, p := range parts {
				n, err := strconv.Atoi(p)
				if err != nil {
					return result, err
				}
				button = append(button, n)
			}
			buttons = append(buttons, button)
		}
		cur.buttons = buttons

		fJoltages := rJoltage.FindStringSubmatch(row)
		parts := strings.Split(fJoltages[1], ",")
		var joltages []int = make([]int, 0)
		for _, j := range parts {
			n, err := strconv.Atoi(j)
			if err != nil {
				return result, err
			}
			joltages = append(joltages, n)
		}
		cur.joltages = joltages

		result = append(result, cur)
	}

	return result, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func toggleLights(lights []bool, mac machine, button int, count int) int {
	found := true
	for i := range lights {
		if lights[i] != mac.lights[i] {
			found = false
			break
		}
	}

	if found {
		return count
	}

	if button >= len(mac.buttons) {
		return math.MaxInt
	}

	lightsClick, lightsNoClick := make([]bool, len(lights)), make([]bool, len(lights))
	copy(lightsClick, lights)
	copy(lightsNoClick, lights)

	for i := 0; i < len(mac.buttons[button]); i++ {
		lightsClick[mac.buttons[button][i]] = !lightsClick[mac.buttons[button][i]]
	}

	return min(toggleLights(lightsNoClick, mac, button+1, count), toggleLights(lightsClick, mac, button+1, count+1))
}

func toggleJoltages(mac machine) (int, error) {
	n, m := len(mac.joltages), len(mac.buttons)
	target := mac.joltages
	var A [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
		for j, button := range mac.buttons {
			for _, c := range button {
				if c == i {
					A[i][j] = 1
				}
			}
		}
	}

	// Create Z3 context
	config := z3.NewContextConfig()
	context := z3.NewContext(config)

	// Create solver
	solver := z3.NewSolver(context)
	intSort := context.IntSort()

	zero := context.FromInt(0, intSort).(z3.Int)

	// Create integer variables x_j >= 0
	x := make([]z3.Int, m)
	for j := 0; j < m; j++ {
		x[j] = context.IntConst(fmt.Sprintf("x_%d", j))
		// x_j >= 0
		solver.Assert(x[j].GE(zero))
	}

	// Constraints: sum_j A[i][j] * x_j == target[i]
	for i := 0; i < n; i++ {
		terms := make([]z3.Int, 0)
		for j := 0; j < m; j++ {
			if A[i][j] == 1 {
				coef := context.FromInt(int64(A[i][j]), intSort).(z3.Int)
				terms = append(terms, coef.Mul(x[j]))
			}
		}

		var sum z3.Int = terms[0].Add(terms[1:]...)

		rhs := context.FromInt(int64(target[i]), intSort).(z3.Int)
		solver.Assert(sum.Eq(rhs))
	}

	// Objective: minimize total presses = sum_j x_j
	t := context.IntConst("total")
	var total z3.Int = x[0].Add(x[1:]...)
	solver.Assert(t.Eq(total))
	solver.Assert(t.GE(zero))

	bestTotal := math.MaxInt

	for {
		satisfied, err := solver.Check()
		if err != nil {
			return bestTotal, err
		}
		if !satisfied {
			break
		}

		model := solver.Model()

		tVal, _, _ := model.Eval(t, true).(z3.Int).AsInt64()
		curTotal := int(tVal)

		if curTotal < bestTotal {
			bestTotal = curTotal
		}

		if curTotal == 0 {
			break
		}

		bound := context.FromInt(int64(curTotal), intSort).(z3.Int)
		solver.Assert(t.LT(bound))
	}

	if bestTotal == math.MaxInt {
		return bestTotal, errors.New("couldn't satisfy")
	}

	return bestTotal, nil
}

func Challenge1(data []string) (int, error) {
	machines, err := parseInput(data)
	if err != nil {
		return 0, err
	}

	var result int = 0

	for _, mac := range machines {
		var lights []bool = make([]bool, len(mac.lights))
		result += toggleLights(lights, mac, 0, 0)
	}

	return result, nil
}

func Challenge2(data []string) (int, error) {
	machines, err := parseInput(data)
	if err != nil {
		return 0, err
	}

	var result int = 0

	for _, mac := range machines {
		tmp, err := toggleJoltages(mac)
		if err != nil {
			// return result, err
			continue
		}
		result += tmp
	}

	return result, nil
}

func Run() {
	fmt.Println("Day 10 - Factory")
	path := "factory/input.txt"
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
