package utils

import "math"

func WilsonScoreFromVotes(upvotes, downvotes int) float64 {
	total := upvotes + downvotes
	if total == 0 {
		return 0
	}

	n := float64(total)
	p := float64(upvotes) / n
	z := 1.96
	zSquared := z * z

	left := p + zSquared/(2*n)
	right := z * math.Sqrt((p*(1-p)+zSquared/(4*n))/n)
	denominator := 1 + zSquared/n

	return (left - right) / denominator
}

func TimeDecay(ageDays float64) float64 {
	gravity := 180.0
	return 1.0 / (1.0 + ageDays/gravity)
}

func CalculateWilsonScoreWithTimeDecay(score, ageDays float64) float64 {
	decayFactor := TimeDecay(ageDays)
	return score * decayFactor
}