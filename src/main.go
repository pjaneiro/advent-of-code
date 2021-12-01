package main

import (
	"fmt"
	"sonarsweep"
	"time"
)

func main() {
	timerAll := time.Now()
	timer := time.Now()

	timer = time.Now()
	sonarsweep.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
