package main

import (
	"fmt"
	"time"

	"github.com/pjaneiro/advent-of-code/2023/cubeconundrum"
	"github.com/pjaneiro/advent-of-code/2023/trebuchet"
)

func main() {
	timerAll := time.Now()
	timer := time.Now()

	timer = time.Now()
	trebuchet.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	cubeconundrum.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
