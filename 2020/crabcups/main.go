package crabcups

import (
	"container/ring"
	"fmt"
	"strconv"
)

func createLinkedList(data []int, size int) []*ring.Ring {
	result := ring.New(size)
	var pointers []*ring.Ring = make([]*ring.Ring, size+1)
	for i := 0; i < len(data); i++ {
		result.Value = data[i]
		pointers[data[i]] = result
		result = result.Next()
	}
	if size > len(data) {
		for i := len(data) + 1; i <= size; i++ {
			result.Value = i
			pointers[i] = result
			result = result.Next()
		}
	}
	return pointers
}

func solve(cups []int, size int, steps int) []*ring.Ring {
	pointers := createLinkedList(cups, size)
	cur := pointers[cups[0]]

	for move := 0; move < steps; move++ {
		toMove := cur.Unlink(3)

		destValue := (cur.Value).(int) - 1
		if destValue < 1 {
			destValue = size
		}
		for destValue == (toMove.Value).(int) || destValue == (toMove.Next().Value).(int) || destValue == (toMove.Prev().Value).(int) {
			destValue = destValue - 1
			if destValue < 1 {
				destValue = size
			}
		}

		dest := pointers[destValue]
		dest.Link(toMove)

		cur = cur.Next()
	}
	return pointers
}

func Challenge1(cups []int) (string, error) {
	pointers := solve(cups, 9, 100)

	result := ""

	for cur := pointers[1].Next(); cur.Value != 1; cur = cur.Next() {
		result = result + strconv.Itoa((cur.Value).(int))
	}

	return result, nil
}

func Challenge2(cups []int) (string, error) {
	pointers := solve(cups, 1000000, 10000000)

	res1, res2 := pointers[1].Next(), pointers[1].Next().Next()
	prod := int64(int64((res1.Value).(int)) * int64((res2.Value).(int)))

	return strconv.FormatInt(prod, 10), nil
}

func Run() {
	fmt.Println("Day 23 - Crab Cups")
	data := []int{4, 5, 9, 6, 7, 2, 8, 1, 3}

	var result string
	var err error
	result, err = Challenge1(data)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %s\n", result)
	}

	result, err = Challenge2(data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %s\n", result)
	}
}
