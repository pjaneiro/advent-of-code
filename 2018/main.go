package main

import (
	"fmt"
	"github.com/pjaneiro/advent-of-code/2018/chronalcalibration"
	"time"
)

func main() {
	timerAll := time.Now()
	timer := time.Now()

	timer = time.Now()
	chronalcalibration.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
