package nospaceleftondevice

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Node struct {
	Name     string
	Parent   *Node
	Children []*Node
	Size     int
}

const INFO = false

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func createNode(name string, parent *Node) *Node {
	node := Node{Name: name, Parent: parent}
	node.Children = make([]*Node, 0)
	return &node
}

func findChild(name string, parent *Node) *Node {
	for _, child := range parent.Children {
		if child.Name == name {
			return child
		}
	}
	return nil
}

func makeTree(data []string, root *Node) error {
	var current *Node = root

	r, err := regexp.Compile(`^(\d+) (.*)$`)
	if err != nil {
		return errors.New("couldn't parse regex expression")
	}

	for _, line := range data {
		if line == "$ ls" {
			continue
		}
		if line == "$ cd /" {
			current = root
			continue
		}
		if line == "$ cd .." {
			current = current.Parent
			continue
		}
		if line[0:4] == "$ cd" {
			if child := findChild(line[5:], current); child != nil {
				current = child
				continue
			}
			newChild := createNode(line[5:], current)
			current.Children = append(current.Children, newChild)
			current = newChild
			continue
		}
		if line[0:3] == "dir" {
			if child := findChild(line[4:], current); child != nil {
				continue
			}
			newChild := createNode(line[4:], current)
			current.Children = append(current.Children, newChild)
			continue
		}
		data := r.FindStringSubmatch(line)
		if child := findChild(data[2], current); child != nil {
			continue
		}
		newChild := createNode(data[2], current)
		if newChild.Size, err = strconv.Atoi(data[1]); err != nil {
			return errors.New("error parsing file size")
		}
		current.Children = append(current.Children, newChild)
		continue
	}

	return nil
}

func calculateWeight(node *Node, cache map[*Node]int) int {
	if len(node.Children) == 0 {
		return node.Size
	}
	if val, ok := cache[node]; ok {
		return val
	}
	result := 0
	for _, each := range node.Children {
		result += calculateWeight(each, cache)
	}
	cache[node] = result
	return result
}

func printTree(node *Node, indentation int) {
	for i := 0; i < indentation; i++ {
		fmt.Printf(" ")
	}
	if len(node.Children) == 0 {
		fmt.Printf("- %s (file, size=%d)\n", node.Name, node.Size)
	} else {
		fmt.Printf("- %s (dir)\n", node.Name)
		for _, each := range node.Children {
			printTree(each, indentation+2)
		}
	}
}

func Challenge1(data []string) (int, error) {
	var weights map[*Node]int = make(map[*Node]int)
	var root *Node = createNode("/", nil)

	if makeTree(data, root) != nil {
		return 0, errors.New("error making tree")
	}

	if INFO {
		printTree(root, 0)
	}

	calculateWeight(root, weights)

	result := 0
	for _, val := range weights {
		if val < 100000 {
			result += val
		}
	}

	return result, nil
}

func Challenge2(data []string) (int, error) {
	const MAX_SPACE = 70000000
	const MIN_SPACE = 30000000

	var weights map[*Node]int = make(map[*Node]int)
	var root *Node = createNode("/", nil)

	if makeTree(data, root) != nil {
		return 0, errors.New("error making tree")
	}

	if INFO {
		printTree(root, 0)
	}

	calculateWeight(root, weights)

	needed := MIN_SPACE - (MAX_SPACE - weights[root])

	difference, result := math.MaxInt, 0
	for _, val := range weights {
		if val < needed {
			continue
		}
		if val-needed < difference {
			difference = val - needed
			result = val
		}
	}

	return result, nil
}

func Run() {
	fmt.Println("Day 7 - No Space Left On Device")
	path := "nospaceleftondevice/input.txt"
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
