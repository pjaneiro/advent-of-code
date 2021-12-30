package main

import (
	"fmt"
	"github.com/pjaneiro/advent-of-code-2019/a1202programalarm"
	"github.com/pjaneiro/advent-of-code-2019/amplificationcircuit"
	"github.com/pjaneiro/advent-of-code-2019/carepackage"
	"github.com/pjaneiro/advent-of-code-2019/crossedwires"
	"github.com/pjaneiro/advent-of-code-2019/monitoringstation"
	"github.com/pjaneiro/advent-of-code-2019/securecontainer"
	"github.com/pjaneiro/advent-of-code-2019/sensorboost"
	"github.com/pjaneiro/advent-of-code-2019/spaceimageformat"
	"github.com/pjaneiro/advent-of-code-2019/spacepolice"
	"github.com/pjaneiro/advent-of-code-2019/spacestoichiometry"
	"github.com/pjaneiro/advent-of-code-2019/sunnywithachanceofasteroids"
	"github.com/pjaneiro/advent-of-code-2019/thenbodyproblem"
	"github.com/pjaneiro/advent-of-code-2019/thetyrannyoftherocketequation"
	"github.com/pjaneiro/advent-of-code-2019/universalorbitmap"
	"time"
)

func main() {
	timerAll := time.Now()
	timer := time.Now()

	timer = time.Now()
	thetyrannyoftherocketequation.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	a1202programalarm.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	crossedwires.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	securecontainer.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	sunnywithachanceofasteroids.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	universalorbitmap.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	amplificationcircuit.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	spaceimageformat.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	sensorboost.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	monitoringstation.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	spacepolice.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	thenbodyproblem.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	carepackage.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	spacestoichiometry.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
