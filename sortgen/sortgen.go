// sortgen holds the implementations of the most common sorting and permutation algorithms.
package sortgen

import (
	"math"
	"math/rand"
	"time"

	"github.com/paulidealiste/goalgs/rangen"
	"github.com/paulidealiste/goalgs/utilgen"
)

// Bubble sort proceeds by traversing the target array and compares each pair of
// adjacent items thus sorting/swapping them if needed.
func Bubblesort(inslice []float64) []float64 {
	defer utilgen.Timetracker(time.Now(), "Bubblesort")
	outslice := make([]float64, len(inslice))
	copy(outslice, inslice)
	for i := 0; i < len(outslice); i++ {
		for j := len(outslice) - 1; j >= i+1; j-- {
			if outslice[j] < outslice[j-1] {
				utilgen.Swapitems(outslice[j-1 : j+1])
			}
		}
	}
	return outslice
}

// Insertsort utilizes the insertion sort algorithm which proceeds by iteration where
// in each iteration step the array element is taken and compared with all of the
// previous elements and insterted in the position when it is found to be less than
// either one of the elements in the target array.
func Insertsort(inslice []float64) []float64 {
	defer utilgen.Timetracker(time.Now(), "Insertsort")
	outslice := make([]float64, len(inslice))
	copy(outslice, inslice)
	for j := 1; j < len(outslice); j++ {
		key := outslice[j]
		i := j - 1
		for i >= 0 && outslice[i] > key {
			outslice[i+1] = outslice[i]
			i = i - 1
		}
		outslice[i+1] = key
	}
	return outslice
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

// Sortpermute performs randomization of input array elements by utilizing sorting
// of the original array elments according to the array of random priorities.
func Sortpermute(inslice []float64) []float64 {
	defer utilgen.Timetracker(time.Now(), "Sortpermute")
	outslice := make([]float64, len(inslice))
	innerpriority := rangen.Gorpa(len(outslice))
	for i, v := range innerpriority {
		outslice[i] = inslice[v]
	}
	return outslice
}

// Inplacepermute randomizes the order of the array elements by swapping randomly
// chosen pairings during one traversing of the original, input array.
func Inplacepermute(inslice []float64) []float64 {
	defer utilgen.Timetracker(time.Now(), "Inplacepermute")
	outslice := make([]float64, len(inslice))
	copy(outslice, inslice)
	source := rand.NewSource(time.Now().UnixNano())
	randomer := rand.New(source)
	for i := len(outslice) - 1; i > 0; i-- {
		j := randomer.Intn(i)
		outslice[i], outslice[j] = outslice[j], outslice[i]
	}
	return outslice
}
