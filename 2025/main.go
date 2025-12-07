package main

import (
	"fmt"
	"time"

	"github.com/pjaneiro/advent-of-code/2025/cafeteria"
	"github.com/pjaneiro/advent-of-code/2025/giftshop"
	"github.com/pjaneiro/advent-of-code/2025/laboratories"
	"github.com/pjaneiro/advent-of-code/2025/lobby"
	"github.com/pjaneiro/advent-of-code/2025/printingdepartment"
	"github.com/pjaneiro/advent-of-code/2025/secretentrance"
	"github.com/pjaneiro/advent-of-code/2025/trashcompactor"
)

func main() {
	timerAll := time.Now()
	var timer time.Time

	timer = time.Now()
	secretentrance.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	giftshop.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	lobby.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	printingdepartment.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	cafeteria.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	trashcompactor.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	laboratories.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
