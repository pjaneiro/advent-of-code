package main

import (
	"fmt"
	"time"

	"github.com/pjaneiro/advent-of-code/2022/caloriecounting"
	"github.com/pjaneiro/advent-of-code/2022/rockpaperscissors"
)

func main() {
	timerAll := time.Now()
	timer := time.Now()

	timer = time.Now()
	caloriecounting.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	rockpaperscissors.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}