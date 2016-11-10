//sortgen holds the implementations of the common sorting algorithms
package sortgen

import (
	"math"
	"testing"
	"time"

	"github.com/paulidealiste/goalgs/utilgen"
)

//Insertsort serves as the insertion sort algorithm working on the slice of float64.
func Insertsort(inslice []float64) []float64 {
	defer utilgen.Timetracker(time.Now(), "Insertsort")
	for j := 1; j < len(inslice); j++ {
		key := inslice[j]
		i := j - 1
		for i >= 0 && inslice[i] > key {
			inslice[i+1] = inslice[i]
			i = i - 1
		}
		inslice[i+1] = key
	}
	return inslice
}

//Mergesort is a dynamic top-level sorting function utilizing divide-and-conquer approach.
func Mergesort(inslice []float64, p int, r int) []float64 {
	//defer utilgen.Timetracker(time.Now(), "Mergesort")
	var outslice []float64
	if p < r {
		q := int(math.Floor(float64((p + r) / 2)))
		Mergesort(inslice, p, q)
		Mergesort(inslice, q, r-1)
		outslice = Innermerge(inslice, p, q, r)
	}
	return outslice
}

func Innermerge(inslice []float64, p int, q int, r int) []float64 {
	n1 := q - p + 1
	n2 := r - q
	innerleft := make([]float64, n1+1)
	innerright := make([]float64, n2+1)
	for i := 0; i < n1; i++ {
		innerleft[i] = inslice[p+i]
	}
	for j := 0; j < n2; j++ {
		innerright[j] = inslice[q+j]
	}
	innerleft[n1] = math.MaxFloat64
	innerright[n2] = math.MaxFloat64
	ii := 0
	jj := 0
	for k := p; k < r; k++ {
		if innerleft[ii] <= innerright[jj] {
			inslice[k] = innerleft[ii]
			ii++
		} else {
			inslice[k] = innerright[jj]
			jj++
		}
	}
	return inslice
}

//Benchmarks and tests
func BenchmarkInsertsort(b *testing.B) {
	ta := []float64{10.0, 8.7, 6.3, 4.2, 9.2, 5.8, 3.1, 2.3, 1.1}
	for i := 0; i < b.N; i++ {
		Insertsort(ta)
	}
}
