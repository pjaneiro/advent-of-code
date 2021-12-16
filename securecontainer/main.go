package securecontainer

import (
	"fmt"
)

func Challenge1(min int, max int) (int, error) {
	result := 0
	for i := min; i <= max; i++ {
		if i/100000 != i/10000%10 && i/10000%10 != i/1000%10 && i/1000%10 != i/100%10 && i/100%10 != i/10%10 && i/10%10 != i%10 {
			continue
		}
		if i/100000 > i/10000%10 || i/10000%10 > i/1000%10 || i/1000%10 > i/100%10 || i/100%10 > i/10%10 || i/10%10 > i%10 {
			continue
		}
		result++
	}
	return result, nil
}

func Challenge2(min int, max int) (int, error) {
	result := 0
	for i := min; i <= max; i++ {
		n1, n2, n3, n4, n5, n6 := i/100000, i/10000%10, i/1000%10, i/100%10, i/10%10, i%10
		if n1 != n2 && n2 != n3 && n3 != n4 && n4 != n5 && n5 != n6 {
			continue
		}
		if n1 > n2 || n2 > n3 || n3 > n4 || n4 > n5 || n5 > n6 {
			continue
		}
		parts := []int{n1, n2, n3, n4, n5, n6}
		cnt := 1
		for i := 1; i < 6; i++ {
			if parts[i] == parts[i-1] {
				cnt++
			} else if cnt == 2 {
				break
			} else {
				cnt = 1
			}
		}
		if cnt != 2 {
			continue
		}
		result++
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 4 - Secure Container")
	min, max := 125730, 579381

	var result int
	result, err := Challenge1(min, max)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(min, max)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
