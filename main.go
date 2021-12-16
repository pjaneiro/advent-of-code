package main

import (
	"fmt"
	"time"
)

func main() {
	timerAll := time.Now()

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
