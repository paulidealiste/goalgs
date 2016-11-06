//sortgen holds the implementations of the common sorting algorithms
package sortgen

import (
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

//Benchmarks and tests
func BenchmarkInsertsort(b *testing.B) {
	ta := []float64{10.0, 8.7, 6.3, 4.2, 9.2, 5.8, 3.1, 2.3, 1.1}
	for i := 0; i < b.N; i++ {
		Insertsort(ta)
	}
}
