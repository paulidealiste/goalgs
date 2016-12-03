// datagen holds procedures for creating various data structures
package datagen

import (
	"time"

	"github.com/paulidealiste/goalgs/utilgen"
)

type Heap struct {
	Heapsize int
	Length   int
	Inslice  []float64
}

// Heapgen generates a max-heap data structure which is a tree-like representation
// of the input array where each value of a parent element is greater than that of
// a child.
func Heapgen(Inslice []float64) Heap {
	defer utilgen.Timetracker(time.Now(), "Heapgen")
	innerslice := make([]float64, len(Inslice))
	copy(innerslice, Inslice)
	inheap := Heap{len(Inslice) - 1, len(Inslice) - 1, innerslice}
	for i := inheap.Length / 2; i >= 0; i-- {
		Maxheapmaintain(&inheap, i)
	}
	return inheap
}

func Maxheapmaintain(inheap *Heap, driver int) {
	l := 2*driver + 1
	r := l + 1
	largest := driver
	if l < inheap.Heapsize && inheap.Inslice[l] > inheap.Inslice[driver] {
		largest = l
	}
	if r <= inheap.Heapsize && inheap.Inslice[r] > inheap.Inslice[largest] {
		largest = r
	}
	if largest != driver {
		inheap.Inslice[driver], inheap.Inslice[largest] = inheap.Inslice[largest], inheap.Inslice[driver]
		Maxheapmaintain(inheap, largest)
	}
}
