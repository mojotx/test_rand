package result

import (
	"crypto/rand"
	"fmt"
	"io"
	"math"
	"math/big"
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
	var n uint64

	var min, max uint64
	min = math.MaxUint64

	var challenge, champion uint64

	for ; n < 1000; n++ {
		ri, _ := rand.Int(rand.Reader, percentage)

		n := ri.Uint64()

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}

		if n < 30 {
			challenge++
		} else {
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

func CheckResult(w io.Writer, n int, r Result) {

	//fmt.Fprintf(w, "\x1b[1;30m%+v\x1b[0m\n", r)
	if r.ChallengePercent <= 20.0 {

		fmt.Fprintf(w, "%d: %d (%f%%) - %d (%f%%) - %d - %d\n", n,
			r.Champion, r.ChampionPercent,
			r.Challenge, r.ChallengePercent,
			r.Min, r.Max)
	}
}
