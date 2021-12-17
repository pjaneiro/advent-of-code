package crabcups

import (
	"fmt"
	"strconv"
)

type Edges struct {
	Prev int
	Next int
}

func createLinkedList(data []int, size int) map[int]Edges {
	var result map[int]Edges = make(map[int]Edges, size)
	for i := 0; i < 9; i++ {
		result[data[i]] = Edges{Prev: data[(i+8)%9], Next: data[(i+10)%9]}
	}
	if size > 9 {
		result[10] = Edges{Prev: data[len(data)-1], Next: 11}
		for i := 11; i < size; i++ {
			result[i] = Edges{Prev: i - 1, Next: i + 1}
		}
		result[1000000] = Edges{Prev: 999999, Next: data[0]}
		result[data[0]] = Edges{Prev: 1000000, Next: result[data[0]].Next}
		result[data[8]] = Edges{Prev: result[data[8]].Prev, Next: 10}
	}
	return result
}

func Challenge1(cups []int) (string, error) {
	list := createLinkedList(cups, 9)

	cur := cups[0]
	for move := 0; move < 100; move++ {
		next1 := list[cur].Next
		next2 := list[next1].Next
		next3 := list[next2].Next

		dest := cur - 1
		if dest < 1 {
			dest = 9
		}
		for dest == next1 || dest == next2 || dest == next3 {
			dest = dest - 1
			if dest < 1 {
				dest = 9
			}
		}

		list[list[next3].Next] = Edges{Prev: cur, Next: list[list[next3].Next].Next}
		list[cur] = Edges{Prev: list[cur].Prev, Next: list[next3].Next}

		list[list[dest].Next] = Edges{Prev: next3, Next: list[list[dest].Next].Next}
		list[next3] = Edges{Prev: list[next3].Prev, Next: list[dest].Next}
		list[next1] = Edges{Prev: dest, Next: next2}
		list[dest] = Edges{Prev: list[dest].Prev, Next: next1}

		cur = list[cur].Next
	}

	result := ""

	for cur := list[1].Next; cur != 1; cur = list[cur].Next {
		result = result + strconv.Itoa(cur)
	}

	return result, nil
}

func Challenge2(cups []int) (string, error) {
	list := createLinkedList(cups, 1000000)

	cur := cups[0]
	for move := 0; move < 10000000; move++ {
		next1 := list[cur].Next
		next2 := list[next1].Next
		next3 := list[next2].Next

		dest := cur - 1
		if dest < 1 {
			dest = 1000000
		}
		for dest == next1 || dest == next2 || dest == next3 {
			dest = dest - 1
			if dest < 1 {
				dest = 1000000
			}
		}

		list[list[next3].Next] = Edges{Prev: cur, Next: list[list[next3].Next].Next}
		list[cur] = Edges{Prev: list[cur].Prev, Next: list[next3].Next}

		list[list[dest].Next] = Edges{Prev: next3, Next: list[list[dest].Next].Next}
		list[next3] = Edges{Prev: list[next3].Prev, Next: list[dest].Next}
		list[next1] = Edges{Prev: dest, Next: next2}
		list[dest] = Edges{Prev: list[dest].Prev, Next: next1}

		cur = list[cur].Next
	}

	res1, res2 := list[1].Next, list[list[1].Next].Next
	prod := int64(int64(res1) * int64(res2))

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
