package distresssignal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Node struct {
	Value    int
	Children []*Node
	Parent   *Node
}

func readLines(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var pair []string
		pair = append(pair, scanner.Text())
		scanner.Scan()
		pair = append(pair, scanner.Text())
		result = append(result, pair)
		scanner.Scan()
	}

	return result, nil
}

func readAsTree(input string) (*Node, error) {
	root := Node{Value: -1, Children: []*Node{}, Parent: nil}
	curNode := &root

	var curNumber string
	for _, char := range input {
		switch char {
		case '[':
			newNode := Node{Value: -1, Children: []*Node{}, Parent: curNode}
			curNode.Children = append(curNode.Children, &newNode)
			curNode = &newNode
		case ']':
			if len(curNumber) > 0 {
				number, err := strconv.Atoi(curNumber)
				if err != nil {
					return nil, errors.New("error parsing input")
				}
				curNode.Value = number
				curNumber = ""
			}
			curNode = curNode.Parent
		case ',':
			if len(curNumber) > 0 {
				number, err := strconv.Atoi(curNumber)
				if err != nil {
					return nil, errors.New("error parsing input")
				}
				curNode.Value = number
				curNumber = ""
			}
			curNode = curNode.Parent
			newNode := Node{Value: -1, Children: []*Node{}, Parent: curNode}
			curNode.Children = append(curNode.Children, &newNode)
			curNode = &newNode
		default:
			curNumber += string(char)
		}
	}

	return &root, nil
}

func compare(left *Node, right *Node) int {
	if len(left.Children) == 0 && len(right.Children) == 0 {
		if left.Value == right.Value {
			return 0
		}
		if left.Value > right.Value {
			return -1
		}
		return 1
	}
	if left.Value >= 0 {
		return compare(&Node{Value: -1, Children: []*Node{left}, Parent: nil}, right)
	}
	if right.Value >= 0 {
		return compare(left, &Node{Value: -1, Children: []*Node{right}, Parent: nil})
	}
	var count int
	for count = 0; count < len(left.Children) && count < len(right.Children); count++ {
		tmp := compare(left.Children[count], right.Children[count])
		if tmp != 0 {
			return tmp
		}
	}
	if count < len(left.Children) {
		return -1
	}
	if count < len(right.Children) {
		return 1
	}
	return 0
}

func Challenge1(data [][]string) (int, error) {
	result := 0
	for i := 0; i < len(data); i++ {
		var left, right *Node
		var err error
		if left, err = readAsTree(data[i][0]); err != nil {
			return 0, err
		}
		if right, err = readAsTree(data[i][1]); err != nil {
			return 0, err
		}
		if compare(left, right) == 1 {
			result += i + 1
		}
	}
	return result, nil
}

func Challenge2(data [][]string) (int, error) {
	result := 1
	var allSlice []*Node = make([]*Node, 0)
	for i := 0; i < len(data); i++ {
		var cur *Node
		var err error
		if cur, err = readAsTree(data[i][0]); err != nil {
			return 0, err
		}
		allSlice = append(allSlice, cur)
		if cur, err = readAsTree(data[i][1]); err != nil {
			return 0, err
		}
		allSlice = append(allSlice, cur)
	}

	var divider1, divider2 *Node
	var err error
	if divider1, err = readAsTree("[[2]]"); err != nil {
		return 0, err
	}
	allSlice = append(allSlice, divider1)
	if divider2, err = readAsTree("[[6]]"); err != nil {
		return 0, err
	}
	allSlice = append(allSlice, divider2)

	sort.Slice(allSlice, func(i, j int) bool { return compare(allSlice[i], allSlice[j]) == 1 })

	for i := 0; i < len(allSlice); i++ {
		if allSlice[i] == divider1 || allSlice[i] == divider2 {
			result *= (i + 1)
		}
	}

	return result, nil
}

func Run() {
	fmt.Println("Day 13 - Distress Signal")
	path := "distresssignal/input.txt"

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
