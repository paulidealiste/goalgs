// rangen is a utility package for creating slices of random numbers.
package rangen

import (
	"math"
	"math/rand"
	"time"
)

// Gorands wraps Go standard library uniform random generators (math.Rand).
func Gorands(fixed int, gauss bool, scaler float64) []float64 {
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
			randslice[i] = math.Abs(randomer.NormFloat64()) * scaler
		} else {
			randslice[i] = math.Abs(randomer.Float64()) * scaler
		}
	}
	return randslice
}

// Gorpa generates l length array of unique random integers in [0, n] range useful for
// creating the arrays of random priorities (wraps math.Rand permutation)
func Gorpa(n int) []int { //TODO: re-implement by hand
	source := rand.NewSource(time.Now().UnixNano())
	randomer := rand.New(source)
	return randomer.Perm(n)
}

// Samplerint returns random integer in [0,n] range with new seed upon each call
func Simplerint(n int) int {
	source := rand.NewSource(time.Now().UnixNano())
	randomer := rand.New(source)
	return randomer.Intn(n)
}

// Boxmullerrands is a simple implementation of Box-Muller algorithm for generating
// normally distributed random numbers.
func Boxmullerrands(fixed int, mean float64, sigma float64) []float64 {
	var randslice []float64
	source := rand.NewSource(time.Now().UnixNano())
	randomer := rand.New(source)
	if fixed == 0 {
		randslice = make([]float64, randomer.Intn(1000))
	} else {
		randslice = make([]float64, fixed)
	}
	u1 := 0.0
	u2 := 0.0
	for i := 0; i < len(randslice); i++ {
		u1 = math.Abs(randomer.Float64())
		u2 = math.Abs(randomer.Float64())
		randslice[i] = boxmullerinner(u1, u2, mean, sigma)
	}
	return randslice
}

func boxmullerinner(u1 float64, u2 float64, mean float64, sigma float64) float64 {
	bmn := math.Sqrt(-2.0*math.Log(u1)) * math.Cos(2.0*math.Pi*u2)
	return mean + sigma*bmn
}
