package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"

	"github.com/montanaflynn/stats"
)

var limit = big.NewInt(100)

type Result struct {
	Champion         uint64
	ChampionPercent  float64
	Challenge        uint64
	ChallengePercent float64
	Min              uint64
	Max              uint64
}

func doWork() Result {
	var n uint64

	var min, max uint64
	min = math.MaxUint64

	var challenge, champion uint64

	for ; n < 1000; n++ {
		ri, _ := rand.Int(rand.Reader, limit)

		n := ri.Uint64()
		// fmt.Print(n)

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}

		if n < 30 {
			// fmt.Println(" challenge")
			challenge++
		} else {
			// fmt.Println(" champion")
			champion++
		}
	}

	sum := challenge + champion
	return Result{
		Champion:         champion,
		ChampionPercent:  float64(champion) / float64(sum) * 100.0,
		Challenge:        challenge,
		ChallengePercent: float64(challenge) / float64(sum) * 100.0,
		Min:              min,
		Max:              max,
	}
}

func main() {

	var values []float64

	for n := 0; n < 1000; n++ {

		result := doWork()
		values = append(values, result.ChallengePercent)
		if result.ChallengePercent <= 20.0 {
			fmt.Printf("%d: %d (%f%%) - %d (%f%%) - %d - %d\n", n, result.Champion, result.ChampionPercent, result.Challenge, result.ChallengePercent, result.Min, result.Max)
		}

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
