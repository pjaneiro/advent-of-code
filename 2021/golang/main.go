package main

import (
	"fmt"
	"github.com/pjaneiro/advent-of-code-2021/amphipod"
	"github.com/pjaneiro/advent-of-code-2021/arithmeticlogicunit"
	"github.com/pjaneiro/advent-of-code-2021/beaconscanner"
	"github.com/pjaneiro/advent-of-code-2021/binarydiagnostic"
	"github.com/pjaneiro/advent-of-code-2021/chiton"
	"github.com/pjaneiro/advent-of-code-2021/diracdice"
	"github.com/pjaneiro/advent-of-code-2021/dive"
	"github.com/pjaneiro/advent-of-code-2021/dumbooctopus"
	"github.com/pjaneiro/advent-of-code-2021/extendedpolymerization"
	"github.com/pjaneiro/advent-of-code-2021/giantsquid"
	"github.com/pjaneiro/advent-of-code-2021/hydrothermalventure"
	"github.com/pjaneiro/advent-of-code-2021/lanternfish"
	"github.com/pjaneiro/advent-of-code-2021/packetdecoder"
	"github.com/pjaneiro/advent-of-code-2021/passagepathing"
	"github.com/pjaneiro/advent-of-code-2021/reactorreboot"
	"github.com/pjaneiro/advent-of-code-2021/seacucumber"
	"github.com/pjaneiro/advent-of-code-2021/sevensegmentsearch"
	"github.com/pjaneiro/advent-of-code-2021/smokebasin"
	"github.com/pjaneiro/advent-of-code-2021/snailfish"
	"github.com/pjaneiro/advent-of-code-2021/sonarsweep"
	"github.com/pjaneiro/advent-of-code-2021/syntaxscoring"
	"github.com/pjaneiro/advent-of-code-2021/thetreacheryofwhales"
	"github.com/pjaneiro/advent-of-code-2021/transparentorigami"
	"github.com/pjaneiro/advent-of-code-2021/trenchmap"
	"github.com/pjaneiro/advent-of-code-2021/trickshot"
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

	timer = time.Now()
	syntaxscoring.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	dumbooctopus.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	passagepathing.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	transparentorigami.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	extendedpolymerization.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	chiton.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	packetdecoder.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	trickshot.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	snailfish.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	beaconscanner.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	trenchmap.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	diracdice.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	reactorreboot.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	amphipod.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	arithmeticlogicunit.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	seacucumber.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
