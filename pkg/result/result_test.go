package result

import (
	"bytes"
	"testing"
)

func TestDoWork(t *testing.T) {
	r := DoWork()
	if r.Champion == 0 {
		t.Error("Champion should not be zero")
	}
	if r.Challenge == 0 {
		t.Error("Challenge should not be zero")
	}
	if r.Min > 0 {
		t.Error("Min should probably be zero")
	}
	if r.Max == 0 {
		t.Error("Max should not be zero")
	}
	if r.ChampionPercent == 0 {
		t.Error("ChampionPercent should not be zero")
	}
	if r.ChallengePercent == 0 {
		t.Error("ChallengePercent should not be zero")
	}
}

func TestCheckResultGood(t *testing.T) {
	// 936: 708 (70.800000%) - 292 (29.200000%) - 0 - 99
	r := Result{
		Champion:         708,
		ChampionPercent:  70.800000,
		Challenge:        292,
		ChallengePercent: 29.200000,
		Min:              0,
		Max:              99,
	}
	var b bytes.Buffer
	CheckResult(&b, 936, r)

	got := b.String()
	want := ""
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestCheckResultBad(t *testing.T) {
	r := Result{
		Champion:         808,
		ChampionPercent:  80.800000,
		Challenge:        192,
		ChallengePercent: 19.200000,
		Min:              0,
		Max:              99,
	}

	var b bytes.Buffer
	CheckResult(&b, 42, r)
	got := b.String()
	want := "42: 808 (80.800000%) - 192 (19.200000%) - 0 - 99\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
