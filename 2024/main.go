package main

import (
	"fmt"
	"time"

	"github.com/pjaneiro/advent-of-code/2024/historianhysteria"
	"github.com/pjaneiro/advent-of-code/2024/mullitover"
	"github.com/pjaneiro/advent-of-code/2024/rednosedreports"
)

func main() {
	timerAll := time.Now()
	var timer time.Time

	timer = time.Now()
	historianhysteria.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	rednosedreports.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	mullitover.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
