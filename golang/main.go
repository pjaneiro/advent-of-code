package main

import (
	"fmt"
	"github.com/pjaneiro/advent-of-code-2021/binarydiagnostic"
	"github.com/pjaneiro/advent-of-code-2021/dive"
	"github.com/pjaneiro/advent-of-code-2021/giantsquid"
	"github.com/pjaneiro/advent-of-code-2021/hydrothermalventure"
	"github.com/pjaneiro/advent-of-code-2021/lanternfish"
	"github.com/pjaneiro/advent-of-code-2021/sevensegmentsearch"
	"github.com/pjaneiro/advent-of-code-2021/smokebasin"
	"github.com/pjaneiro/advent-of-code-2021/sonarsweep"
	"github.com/pjaneiro/advent-of-code-2021/thetreacheryofwhales"
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

	timer = time.Now()
	binarydiagnostic.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	giantsquid.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	hydrothermalventure.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	lanternfish.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	thetreacheryofwhales.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	sevensegmentsearch.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	smokebasin.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
