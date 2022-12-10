package main

import (
	"fmt"
	"time"

	"github.com/pjaneiro/advent-of-code/2022/caloriecounting"
	"github.com/pjaneiro/advent-of-code/2022/campcleanup"
	"github.com/pjaneiro/advent-of-code/2022/nospaceleftondevice"
	"github.com/pjaneiro/advent-of-code/2022/rockpaperscissors"
	"github.com/pjaneiro/advent-of-code/2022/ropebridge"
	"github.com/pjaneiro/advent-of-code/2022/rucksackreorganization"
	"github.com/pjaneiro/advent-of-code/2022/supplystacks"
	"github.com/pjaneiro/advent-of-code/2022/treetoptreehouse"
	"github.com/pjaneiro/advent-of-code/2022/tuningtrouble"
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

	timer = time.Now()
	rucksackreorganization.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	campcleanup.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	supplystacks.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	tuningtrouble.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	nospaceleftondevice.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	treetoptreehouse.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	ropebridge.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
