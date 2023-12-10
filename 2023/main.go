package main

import (
	"fmt"
	"time"

	"github.com/pjaneiro/advent-of-code/2023/camelcards"
	"github.com/pjaneiro/advent-of-code/2023/cubeconundrum"
	"github.com/pjaneiro/advent-of-code/2023/gearratios"
	"github.com/pjaneiro/advent-of-code/2023/ifyougiveaseedafertilizer"
	"github.com/pjaneiro/advent-of-code/2023/scratchcards"
	"github.com/pjaneiro/advent-of-code/2023/trebuchet"
	"github.com/pjaneiro/advent-of-code/2023/waitforit"
)

func main() {
	timerAll := time.Now()
	var timer time.Time

	timer = time.Now()
	trebuchet.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	cubeconundrum.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	gearratios.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	scratchcards.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	ifyougiveaseedafertilizer.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	waitforit.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	camelcards.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
