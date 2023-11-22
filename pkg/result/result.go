package result

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"

	"github.com/montanaflynn/stats"
)

type Result struct {
	Champion         uint64
	ChampionPercent  float64
	Challenge        uint64
	ChallengePercent float64
	Min              uint64
	Max              uint64
}

var percentage *big.Int

func init() {
	percentage = big.NewInt(100)
}

func DoWork() Result {

	var challenge, champion uint64

	var population []float64

	for i := 0; i < 1000; i++ {

		// Get random integer between 0 and 99
		ri, _ := rand.Int(rand.Reader, percentage)

		n := uint(ri.Uint64())

		population = append(population, float64(n))

		if n < 30 {
			challenge++
		} else {
			champion++
		}
	}

	sum := challenge + champion

	min, err := stats.Min(population)
	if err != nil {
		panic(err)
	}

	max, err := stats.Max(population)
	if err != nil {
		panic(err)
	}

	return Result{
		Champion:         champion,
		ChampionPercent:  float64(champion) / float64(sum) * 100.0,
		Challenge:        challenge,
		ChallengePercent: float64(challenge) / float64(sum) * 100.0,
		Min:              uint64(min),
		Max:              uint64(max),
	}
}

func CheckResult(w io.Writer, n int, r Result) {

	//fmt.Fprintf(w, "\x1b[1;30m%+v\x1b[0m\n", r)
	if r.ChallengePercent <= 20.0 {

		fmt.Fprintf(w, "%d: %d (%f%%) - %d (%f%%) - %d - %d\n", n,
			r.Champion, r.ChampionPercent,
			r.Challenge, r.ChallengePercent,
			r.Min, r.Max)
	}
}
