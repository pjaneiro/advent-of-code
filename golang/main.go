package main

import (
	"fmt"
	"github.com/pjaneiro/advent-of-code-2021/dive"
	"github.com/pjaneiro/advent-of-code-2021/sonarsweep"
	"time"
)

func main() {
	timerAll := time.Now()
	timer := time.Now()

	timer = time.Now()
	sonarsweep.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	dive.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
