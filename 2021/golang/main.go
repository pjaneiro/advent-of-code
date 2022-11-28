package main

import (
	"fmt"
	"github.com/pjaneiro/advent-of-code/2021/golang/amphipod"
	"github.com/pjaneiro/advent-of-code/2021/golang/arithmeticlogicunit"
	"github.com/pjaneiro/advent-of-code/2021/golang/beaconscanner"
	"github.com/pjaneiro/advent-of-code/2021/golang/binarydiagnostic"
	"github.com/pjaneiro/advent-of-code/2021/golang/chiton"
	"github.com/pjaneiro/advent-of-code/2021/golang/diracdice"
	"github.com/pjaneiro/advent-of-code/2021/golang/dive"
	"github.com/pjaneiro/advent-of-code/2021/golang/dumbooctopus"
	"github.com/pjaneiro/advent-of-code/2021/golang/extendedpolymerization"
	"github.com/pjaneiro/advent-of-code/2021/golang/giantsquid"
	"github.com/pjaneiro/advent-of-code/2021/golang/hydrothermalventure"
	"github.com/pjaneiro/advent-of-code/2021/golang/lanternfish"
	"github.com/pjaneiro/advent-of-code/2021/golang/packetdecoder"
	"github.com/pjaneiro/advent-of-code/2021/golang/passagepathing"
	"github.com/pjaneiro/advent-of-code/2021/golang/reactorreboot"
	"github.com/pjaneiro/advent-of-code/2021/golang/seacucumber"
	"github.com/pjaneiro/advent-of-code/2021/golang/sevensegmentsearch"
	"github.com/pjaneiro/advent-of-code/2021/golang/smokebasin"
	"github.com/pjaneiro/advent-of-code/2021/golang/snailfish"
	"github.com/pjaneiro/advent-of-code/2021/golang/sonarsweep"
	"github.com/pjaneiro/advent-of-code/2021/golang/syntaxscoring"
	"github.com/pjaneiro/advent-of-code/2021/golang/thetreacheryofwhales"
	"github.com/pjaneiro/advent-of-code/2021/golang/transparentorigami"
	"github.com/pjaneiro/advent-of-code/2021/golang/trenchmap"
	"github.com/pjaneiro/advent-of-code/2021/golang/trickshot"
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
