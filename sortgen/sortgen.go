// sortgen holds the implementations of the most common sorting and permutation algorithms.
package sortgen

import (
	"math/rand"
	"time"

	"github.com/paulidealiste/goalgs/datagen"
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
// are not merged back but all in sorted order (http://austingwalters.com/merge-sort-in-go-golang/).
func Mergesort(inslice []float64) []float64 {
	defer utilgen.Timetracker(time.Now(), "Mergesort")
	innerslice := make([]float64, len(inslice))
	copy(innerslice, inslice)
	outslice := msrunner(innerslice)
	return outslice
}

func msrunner(inslice []float64) []float64 {
	if len(inslice) < 2 {
		return inslice
	}
	l, r := mergesplit(inslice)
	return innermerge(msrunner(l), msrunner(r))
}

func mergesplit(inslice []float64) ([]float64, []float64) {
	q := len(inslice) / 2
	return inslice[:q], inslice[q:]
}

func innermerge(innerleft, innerright []float64) []float64 {
	s, l, r := len(innerleft)+len(innerright), 0, 0
	innerslice := make([]float64, s, s)
	for n := 0; n < s; n++ {
		if l > len(innerleft)-1 && r <= len(innerright)-1 {
			innerslice[n] = innerright[r]
			r++
		} else if r > len(innerright)-1 && l <= len(innerleft)-1 {
			innerslice[n] = innerleft[l]
			l++
		} else if innerleft[l] > innerright[r] {
			innerslice[n] = innerright[r]
			r++
		} else {
			innerslice[n] = innerleft[l]
			l++
		}
	}
	return innerslice
}

// Heapsort uses the max-heap data structure and proceeds from the root node of the heap
// tree, which holds the largest element, while subsequently decreasing the heap.size
// property leaving only ever smaller sub-max-heaps, until all the elements appear in the
// sorted order reflecting the max-heap structure where each parent is larger that either
// of its children.
func Heapsort(inslice []float64) []float64 {
	defer utilgen.Timetracker(time.Now(), "Heapsort")
	innerslice := make([]float64, len(inslice))
	copy(innerslice, inslice)
	iheap := datagen.Heapgen(innerslice)
	for i := iheap.Length; i >= 1; i-- {
		iheap.Inslice[i], iheap.Inslice[0] = iheap.Inslice[0], iheap.Inslice[i]
		iheap.Heapsize--
		datagen.Maxheapmaintain(&iheap, 0)
	}
	return iheap.Inslice
}

// Quicksort operates in a manner similar to mergesort but with the specific technique
// used for partitioning the array during the divide step. Partitioning is based on the
// selection of the pivot element around wich the partitioning takes place, i.e. all
// elements smaller than the pivot are being moved to the left side of the pivot
// element (http://stackoverflow.com/questions/15802890/idiomatic-quicksort-in-go).
func Quicksort(inslice []float64) []float64 {
	defer utilgen.Timetracker(time.Now(), "Quicksort")
	outslice := make([]float64, len(inslice))
	copy(outslice, inslice)
	quicksortinner(outslice)
	return outslice
}

func quicksortinner(inslice []float64) []float64 {
	if len(inslice) < 2 {
		return inslice
	}
	q := quickpartition(inslice)
	quicksortinner(inslice[:q])
	quicksortinner(inslice[q+1:])
	return inslice
}

func quickpartition(inslice []float64) int {
	pivot := rand.Int() % len(inslice)
	l, r := 0, len(inslice)-1
	inslice[pivot], inslice[r] = inslice[r], inslice[pivot]
	for n := range inslice {
		if inslice[n] < inslice[r] {
			inslice[n], inslice[l] = inslice[l], inslice[n]
			l++
		}
	}
	inslice[l], inslice[r] = inslice[r], inslice[l]
	return l
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
