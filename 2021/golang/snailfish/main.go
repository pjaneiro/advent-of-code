package snailfish

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Node struct {
	Left *Node
	Right *Node
	Value int
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, nil
}

func parseLine(data string) (*Node) {
	rLiteral, err := regexp.Compile(`^\d+$`)
	if err != nil {
		return nil
	}

	if rLiteral.MatchString(data) {
		i, err := strconv.Atoi(data)
		if err != nil {
			return nil
		}
		return &Node{Value: i}
	}

	open, close := 0, 0
	for i := 1; i < len(data) - 1; i++ {
		if data[i] == '[' {
			open++
		} else if data[i] == ']' {
			close++
		} else if data[i] == ',' && open == close {
			return &Node{Left: parseLine(data[1:i]), Right: parseLine(data[i+1:len(data) - 1])}
		}
	}

	return nil
}

func addLeftMost(node *Node, val int) (bool){
	if node.Left == nil && node.Right == nil {
		node.Value = node.Value + val
		return true
	}
	if addLeftMost(node.Left, val) {
		return true
	}
	return addLeftMost(node.Right, val)
}

func addRightMost(node *Node, val int) (bool){
	if node.Left == nil && node.Right == nil {
		node.Value = node.Value + val
		return true
	}
	if addRightMost(node.Right, val) {
		return true
	}
	return addRightMost(node.Left, val)
}

func checkExplodes(node *Node, level int) (bool, int, int) {
	if level >= 4 && node.Left != nil && node.Right != nil {
		l, r := node.Left.Value, node.Right.Value
		node.Left = nil
		node.Right = nil
		node.Value = 0
		return true, l, r
	}
	if node.Left != nil && node.Right != nil {
		expL, lL, rL := checkExplodes(node.Left, level + 1)
		if expL {
			if rL != -1 {
				if addLeftMost(node.Right, rL) {
					return true, lL, -1
				}
			}
			return true, lL, rL
		}
		expR, lR, rR := checkExplodes(node.Right, level + 1)
		if expR {
			if lR != -1 {
				if addRightMost(node.Left, lR) {
					return true, -1, rR
				}
			}
			return true, lR, rR
		}
	}
	return false, -1, -1
}

func checkSplits(node *Node) (bool) {
	if node.Left == nil && node.Right == nil {
		if node.Value >= 10 {
			leftVal := int(math.Floor(float64(node.Value) / 2.0))
			rightVal := int(math.Ceil(float64(node.Value) / 2.0))
			node.Left = &Node{Value: leftVal}
			node.Right = &Node{Value: rightVal}
			node.Value = 0
			return true
		}
	} else {
		if checkSplits(node.Left) {
			return true
		}
		return checkSplits(node.Right)
	}
	return false
}

func cloneTree(node *Node) (*Node) {
	if node.Left == nil && node.Right == nil {
		return &Node{Value: node.Value}
	}
	return &Node{Left: cloneTree(node.Left), Right: cloneTree(node.Right)}
}

func reduceTree(node *Node) (*Node) {
	result := cloneTree(node)
	for true {
		if ok, _, _ := checkExplodes(result, 0); ok {
			continue
		}
		if checkSplits(result) {
			continue
		}
		break
	}
	return result
}

func addTrees(left *Node, right *Node) (*Node) {
	return &Node{Left: left, Right: right}
}

func printTree(node *Node, level int) {
	if node.Left == nil && node.Right == nil {
		for i := 0; i < level; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("%d\n", node.Value)
	} else {
		printTree(node.Left, level+1)
		printTree(node.Right, level+1)
	}
}

func treeSum(node *Node) (int) {
	if node.Left == nil && node.Right == nil {
		return node.Value
	}
	return treeSum(node.Left) + treeSum(node.Right)
}

func treeMagnitude(node *Node) (int) {
	if node.Left == nil && node.Right == nil {
		return node.Value
	}
	return (3 * treeMagnitude(node.Left)) + (2 * treeMagnitude(node.Right))
}

func Challenge1(data []string) (int, error) {
	var finalTree *Node = reduceTree(parseLine(data[0]))
	for i := 1; i < len(data); i++ {
		finalTree = reduceTree(addTrees(finalTree, reduceTree(parseLine(data[i]))))
	}
	return treeMagnitude(finalTree), nil
}

func Challenge2(data []string) (int, error) {
	var parsedTrees []*Node = make([]*Node, len(data))
	for i, j := range data {
		parsedTrees[i] = reduceTree(parseLine(j))
	}
	best := 0
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data); j++ {
			if i == j {
				continue
			}
			result := treeMagnitude(reduceTree(addTrees(parsedTrees[i], parsedTrees[j])))
			if result > best {
				best = result
			}
		}
	}
	return best, nil
}

func Run() {
	fmt.Println("Day 18 - Snailfish")
	path := "snailfish/input.txt"
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
