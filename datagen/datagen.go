// datagen holds procedures for creating various data structures
package datagen

import (
	"errors"
	"math"
	"time"

	"github.com/paulidealiste/goalgs/utilgen"
)

type Heap struct {
	Heapsize int
	Length   int
	Inslice  []float64
}

type Bst struct {
	Bstsize    int
	Rootnode   *Bstnode
	Innerslice []float64
	Inslice    []Bstnode
}

type Bstnode struct {
	data       float64
	leftchild  *Bstnode
	rightchild *Bstnode
	parent     *Bstnode
}

// Bstgen tries to transform input slice into a binary search tree representation
// where each datum is reperesented by a node with a key which satisfies criteria
// that key in left child is always smaller than the key of the parent while the
// same key in the right child is always larger.
func Bstgen(inslice []float64) Bst {
	defer utilgen.Timetracker(time.Now(), "Bstgen")
	innerslice := make([]float64, len(inslice))
	copy(innerslice, inslice)
	var bstlist []Bstnode
	innerbst := Bst{0, nil, innerslice, bstlist}
	bstgencreate(&innerbst)
	return innerbst
}

func bstgencreate(inbst *Bst) {
	for i := 0; i < len(inbst.Innerslice)-1; i++ {
		insertnode := Bstnode{inbst.Innerslice[i], nil, nil, nil}
		inbst.Insert(&insertnode)
	}
}

// Binary search tree maintenance methods

// Bst.Insert() acts by inserting new Bstnode element at appropriate position
// so that basic premises of the binary search tree are maintained
func (inbst *Bst) Insert(insertnode *Bstnode) {
	if inbst.Rootnode == nil {
		inbst.Rootnode = insertnode
	} else {
		currentRoot := inbst.Rootnode
		for {
			if insertnode.data > currentRoot.data {
				if currentRoot.rightchild == nil {
					currentRoot.rightchild = insertnode
					inbst.Inslice = append(inbst.Inslice, *currentRoot)
					inbst.Bstsize++
					break
				} else {
					currentRoot = currentRoot.rightchild
				}
			} else {
				if currentRoot.leftchild == nil {
					currentRoot.leftchild = insertnode
					inbst.Inslice = append(inbst.Inslice, *currentRoot)
					inbst.Bstsize++
					break
				} else {
					currentRoot = currentRoot.leftchild
				}
			}
		}
	}
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

// Heap maintenance methods // assuming that keys are indices to a Heap.Inslice slice

// Heap.maximum() returns the maximum of the generated max-heap
func (h *Heap) Maximum() float64 {
	return h.Inslice[0]
}

// Heap.extractmax() removes and returns the element of the heap with the largest key
func (h *Heap) Extractmax() (float64, error) {
	if h.Heapsize < 1 {
		return -1.0, errors.New("Heap underflow error. Heap size less than 1.")
	}
	cmax := h.Inslice[0]
	h.Inslice[0] = h.Inslice[h.Heapsize]
	h.Heapsize--
	h.Inslice = h.Inslice[:h.Heapsize+1]
	Maxheapmaintain(h, 0)
	return cmax, nil
}

// Heap.increasekey() goes on to increase the value of the element's current key
// (elementkey) to the new position specified by targetkey's value.
func (h *Heap) Increasekey(elementkey int, targetkey float64) error {
	if targetkey < h.Inslice[elementkey] {
		return errors.New("New key is smaller than the current key, which is not allowed in max-heap increasekey.")
	}
	h.Inslice[elementkey] = targetkey
	for elementkey > 1 && h.Inslice[elementkey>>1] < h.Inslice[elementkey] {
		h.Inslice[elementkey], h.Inslice[elementkey>>1] = h.Inslice[elementkey>>1], h.Inslice[elementkey]
		elementkey = elementkey >> 1
	}
	return nil
}

// Heap.insert() inserts the targetkey value in the appropriate place in the heap
// tree structure thus maintaining the max-heap property
func (h *Heap) Insert(targetkey float64) {
	h.Heapsize++
	h.Inslice = append(h.Inslice, math.MaxFloat64*-1)
	h.Increasekey(h.Heapsize, targetkey)
}
