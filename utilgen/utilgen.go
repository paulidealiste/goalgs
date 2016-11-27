// utilgen holds various utility function primarily for documenting the algs output and performance
// as well as some common generic-like functions
package utilgen

import (
	"errors"
	"fmt"
	"time"
)

// Simple timetracker function called with defer at the onset of the function.
func Timetracker(start time.Time, fname string) {
	elapsed := time.Since(start)
	fmt.Printf("Function %s ran for %s\n", fname, elapsed)
}

//Sum of the supplied values
func Sumfun(inpart []float64) float64 {
	if len(inpart) < 1 {
		err := errors.New("Input array must be at least two elements long.")
		panic(err)
	}
	var outsum float64
	for _, v := range inpart {
		outsum = outsum + v
	}
	return outsum
}

// Return indices of the elements in the supplied array
func Retind(inslice []float64, elems []float64) []int {
	if len(inslice) < 1 || len(elems) != 2 {
		err := errors.New("Array of more than one element is required for inslice while a tuple is required for elems.")
		panic(err)
	}
	var inind []int
	for i, v := range inslice {
		if v == elems[0] || v == elems[1] {
			inind = append(inind, i)
		}
	}
	if inind[1] > inind[0] {
		inind[0], inind[1] = inind[1], inind[0]
	}
	return inind
}

// Swap items in the supplied slice/tuple which sould be a pair of values.
func Swapitems(intuple []float64) []float64 {
	if len(intuple) != 2 {
		err := errors.New("Tuple (slice of length 2) is required for swapping.")
		panic(err)
	}
	intuple[0], intuple[1] = intuple[1], intuple[0]
	return intuple
}
