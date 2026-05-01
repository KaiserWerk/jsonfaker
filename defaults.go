package jsonfaker

import (
	"math/rand"
	"time"
)

func defaultNumberFaker(field string, schema map[string]any, r *rand.Rand) float64 {

	min := 0.0
	max := 100.0

	if v, ok := schema["minimum"].(float64); ok {
		min = v
	}
	if v, ok := schema["maximum"].(float64); ok {
		max = v
	}

	return min + r.Float64()*(max-min)
}

func defaultIntegerFaker(field string, schema map[string]any, r *rand.Rand) int {

	min := 0
	max := 100

	if v, ok := schema["minimum"].(float64); ok {
		min = int(v)
	}
	if v, ok := schema["maximum"].(float64); ok {
		max = int(v)
	}

	if max <= min {
		return min
	}

	return r.Intn(max-min+1) + min
}

func randomDate(r *rand.Rand) time.Time {
	start := time.Now().AddDate(-2, 0, 0).Unix()
	end := time.Now().AddDate(2, 0, 0).Unix()

	sec := r.Int63n(end-start) + start
	return time.Unix(sec, 0)
}
