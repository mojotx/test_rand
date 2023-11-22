package main

import (
	"fmt"
	"math/big"
	"os"

	resultPkg "github.com/mojotx/test_rand/pkg/result"
	"github.com/montanaflynn/stats"
)

var limit = big.NewInt(100)

func main() {

	var values []float64

	for n := 0; n < 1000; n++ {

		result := resultPkg.DoWork()
		values = append(values, result.ChallengePercent)
		resultPkg.CheckResult(os.Stderr, n, result)
	}

	// mean or average
	if mean, err := stats.Mean(values); err != nil {
		panic(err)
	} else {
		fmt.Printf("Mean: %f\n", mean)
	}

	// standard deviation
	if stdDev, err := stats.StandardDeviationPopulation(values); err != nil {
		panic(err)
	} else {
		fmt.Printf("Standard Deviation (σ): %f\n", stdDev)
	}

	// min
	if min, err := stats.Min(values); err != nil {
		panic(err)
	} else {
		fmt.Printf("Min: %f\n", min)
	}

	// max
	if max, err := stats.Max(values); err != nil {
		panic(err)
	} else {
		fmt.Printf("Max: %f\n", max)
	}
}
