// Package rangen is a utility package for creating slices of random numbers.
package rangen

import (
	"math"
	"math/rand"
	"time"
)

//Gorands utilizes the Go standard library uniform random generators.
func Gorands(fixed int, gauss bool) []float64 {
	var randslice []float64
	source := rand.NewSource(time.Now().UnixNano())
	randomer := rand.New(source)
	if fixed == 0 {
		randslice = make([]float64, randomer.Intn(1000))
	} else {
		randslice = make([]float64, fixed)
	}
	for i := 0; i < len(randslice); i++ {
		if gauss == true {
			randslice[i] = math.Abs(randomer.NormFloat64())
		} else {
			randslice[i] = math.Abs(randomer.Float64())
		}
	}
	return randslice
}
