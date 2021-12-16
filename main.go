package main

import (
	"fmt"
	"github.com/pjaneiro/advent-of-code-2020/adapterarray"
	"github.com/pjaneiro/advent-of-code-2020/allergenassessment"
	"github.com/pjaneiro/advent-of-code-2020/binaryboarding"
	"github.com/pjaneiro/advent-of-code-2020/conwaycubes"
	"github.com/pjaneiro/advent-of-code-2020/crabcombat"
	"github.com/pjaneiro/advent-of-code-2020/crabcups"
	"github.com/pjaneiro/advent-of-code-2020/customcustoms"
	"github.com/pjaneiro/advent-of-code-2020/dockingdata"
	"github.com/pjaneiro/advent-of-code-2020/encodingerror"
	"github.com/pjaneiro/advent-of-code-2020/handheldhalting"
	"github.com/pjaneiro/advent-of-code-2020/handyhaversacks"
	"github.com/pjaneiro/advent-of-code-2020/jurassicjigsaw"
	"github.com/pjaneiro/advent-of-code-2020/monstermessages"
	"github.com/pjaneiro/advent-of-code-2020/operationorder"
	"github.com/pjaneiro/advent-of-code-2020/passportprocessing"
	"github.com/pjaneiro/advent-of-code-2020/passwordphilosophy"
	"github.com/pjaneiro/advent-of-code-2020/rainrisk"
	"github.com/pjaneiro/advent-of-code-2020/rambunctiousrecitation"
	"github.com/pjaneiro/advent-of-code-2020/reportrepair"
	"github.com/pjaneiro/advent-of-code-2020/seatingsystem"
	"github.com/pjaneiro/advent-of-code-2020/shuttlesearch"
	"github.com/pjaneiro/advent-of-code-2020/tickettranslation"
	"github.com/pjaneiro/advent-of-code-2020/toboggantrajectory"
	"time"
)

func main() {
	timerAll := time.Now()
	timer := time.Now()

	timer = time.Now()
	reportrepair.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	passwordphilosophy.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	toboggantrajectory.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	passportprocessing.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	binaryboarding.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	customcustoms.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	handyhaversacks.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	handheldhalting.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	encodingerror.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	adapterarray.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	seatingsystem.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	rainrisk.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	shuttlesearch.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	dockingdata.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	rambunctiousrecitation.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	tickettranslation.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	conwaycubes.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	operationorder.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	monstermessages.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	jurassicjigsaw.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	allergenassessment.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	crabcombat.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	timer = time.Now()
	crabcups.Run()
	fmt.Printf("Time elapsed: %v\n\n", time.Since(timer))

	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
