package main

import (
	"fmt"
	"time"

	"github.com/pjaneiro/advent-of-code/2018/alchemicalreduction"
	"github.com/pjaneiro/advent-of-code/2018/chronalcalibration"
	"github.com/pjaneiro/advent-of-code/2018/chronalcoordinates"
	"github.com/pjaneiro/advent-of-code/2018/inventorymanagementsystem"
	"github.com/pjaneiro/advent-of-code/2018/nomatterhowyousliceit"
	"github.com/pjaneiro/advent-of-code/2018/reposerecord"
)

func main() {
	timerAll := time.Now()
	timer := time.Now()

	timer = time.Now()
	chronalcalibration.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	inventorymanagementsystem.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	nomatterhowyousliceit.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	reposerecord.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	alchemicalreduction.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	chronalcoordinates.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
