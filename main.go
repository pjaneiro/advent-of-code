package main

import (
	"fmt"
	"github.com/pjaneiro/advent-of-code-2020/adapterarray"
	"github.com/pjaneiro/advent-of-code-2020/binaryboarding"
	"github.com/pjaneiro/advent-of-code-2020/conwaycubes"
	"github.com/pjaneiro/advent-of-code-2020/customcustoms"
	"github.com/pjaneiro/advent-of-code-2020/dockingdata"
	"github.com/pjaneiro/advent-of-code-2020/encodingerror"
	"github.com/pjaneiro/advent-of-code-2020/handheldhalting"
	"github.com/pjaneiro/advent-of-code-2020/handyhaversacks"
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
	elapsed := time.Since(timer)

	reportrepair.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	passwordphilosophy.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	toboggantrajectory.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	passportprocessing.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	binaryboarding.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	customcustoms.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	handyhaversacks.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	handheldhalting.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	encodingerror.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	adapterarray.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	seatingsystem.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	rainrisk.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	shuttlesearch.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	dockingdata.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	rambunctiousrecitation.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	tickettranslation.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	conwaycubes.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)

	operationorder.Run()
	elapsed = time.Since(timer)
	timer = time.Now()
	fmt.Printf("Time elapsed: %v\n\n", elapsed)
	fmt.Printf("Total time elapsed: %v\n", time.Since(timerAll))
}
