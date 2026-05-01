package jsonfaker

import (
	"math/rand"
	"time"
)

type Options struct {
	StringFaker  func(field string, schema map[string]any, r *rand.Rand) string
	NumberFaker  func(field string, schema map[string]any, r *rand.Rand) float64
	IntegerFaker func(field string, schema map[string]any, r *rand.Rand) int
	BoolFaker    func(field string, schema map[string]any, r *rand.Rand) bool

	Rand *rand.Rand
}

type Option func(*Options)

func defaultOptions() Options {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return Options{
		Rand: r,

		StringFaker:  defaultStringFaker,
		NumberFaker:  defaultNumberFaker,
		IntegerFaker: defaultIntegerFaker,
		BoolFaker: func(string, map[string]any, *rand.Rand) bool {
			return r.Intn(2) == 1
		},
	}
}

// Optional: deterministischer Seed
func WithSeed(seed int64) Option {
	return func(o *Options) {
		o.Rand = rand.New(rand.NewSource(seed))
	}
}