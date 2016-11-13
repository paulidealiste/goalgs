// sortgen holds the implementations of the most common sorting algorithms.
package sortgen

import (
	"math"
	"testing"
	"time"

	"github.com/paulidealiste/goalgs/utilgen"
)

// Bubble sort proceeds by traversing the target array and compares each pair of
// adjacent items thus sorting/swapping them if needed.
func Bubblesort(inslice []float64) []float64 {
	defer utilgen.Timetracker(time.Now(), "Bubblesort")
	for i := 0; i < len(inslice); i++ {
		for j := len(inslice) - 1; j >= i+1; j-- {
			if inslice[j] < inslice[j-1] {
				utilgen.Swapitems(inslice[j-1 : j+1])
			}
		}
	}
	return inslice
}

// Insertsort utilizes the insertion sort algorithm which proceeds by iteration where
// in each iteration step the array element is taken and compared with all of the
// previous elements and insterted in the position when it is found to be less than
// either one of the elements in the target array.
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

// Mergesort is a dynamic top-level sorting function utilizing divide-and-conquer approach
// where target array is recursively divided to its smallest parts (subarrays arrays of length
// one) and then combined/merged to an ever-longer combined array until all of the subarrays
// are not merged back but all in sorted order.
func Mergesort(inslice []float64, p int, r int) []float64 {
	//defer utilgen.Timetracker(time.Now(), "Mergesort")
	var outslice []float64
	if p < r {
		q := int(math.Floor(float64((p + r) / 2)))
		Mergesort(inslice, p, q)
		Mergesort(inslice, q, r-1)
		outslice = innermerge(inslice, p, q, r)
	}
	return outslice
}

func innermerge(inslice []float64, p int, q int, r int) []float64 {
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

// Benchmarks and tests

func BenchmarkInsertsort(b *testing.B) {
	ta := []float64{10.0, 8.7, 6.3, 4.2, 9.2, 5.8, 3.1, 2.3, 1.1}
	for i := 0; i < b.N; i++ {
		Insertsort(ta)
	}
}

func BenchmarkMergesort(b *testing.B) {
	ta := []float64{10.0, 8.7, 6.3, 4.2, 9.2, 5.8, 3.1, 2.3, 1.1}
	for i := 0; i < b.N; i++ {
		Mergesort(ta, 0, len(ta))
	}
}

func BenchmarkBubblesort(b *testing.B) {
	ta := []float64{10.0, 8.7, 6.3, 4.2, 9.2, 5.8, 3.1, 2.3, 1.1}
	for i := 0; i < b.N; i++ {
		Bubblesort(ta)
	}
}
