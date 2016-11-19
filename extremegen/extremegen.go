// extremegen holds implementations of various extreme findining algorithms
package extremegen

import (
	"math"
	"testing"
	"time"

	"github.com/paulidealiste/goalgs/utilgen"
)

type Localextreme struct {
	index    int
	value    float64
	min, max bool
}

// Maximumsubarray represents the algorithm for finding the nonempty subarray
// of a target array whose values have the largest sum.
func Maximumsubarray(inslice []float64) []float64 {

	return inslice
}

// Onelocalextreme finds either/or one of called for of local extremes and returns the corresponding
// array of structs representing original index, the value of the detected local extreme and the flag
// indicating the type of the extreme.
func Onelocalextreme(inslice []float64) []Localextreme {
	defer utilgen.Timetracker(time.Now(), "Onelocalextreme")
	var outslice []Localextreme
	outslice = Onelocalextrememin(inslice, 0, len(inslice)-1, outslice)
	outslice = Onelocalextrememax(inslice, 0, len(inslice)-1, outslice)
	return outslice
}

func Onelocalextrememin(inslice []float64, p int, r int, outslice []Localextreme) []Localextreme {
	q := int(math.Floor(float64((p + r) / 2)))
	if p <= r && q != r-1 && q > 1 {
		if inslice[q] <= inslice[q-1] && inslice[q] <= inslice[q+1] {
			locex := Localextreme{q, inslice[q], true, false}
			outslice = append(outslice, locex)
		}
		if inslice[q-1] >= inslice[q] && inslice[q] >= inslice[q+1] {
			Onelocalextrememin(inslice, q, r, outslice)
		}
		if inslice[q+1] >= inslice[q] && inslice[q] >= inslice[q-1] {
			Onelocalextrememin(inslice, p, q, outslice)
		}
		if inslice[q-1] <= inslice[q] && inslice[q+1] <= inslice[q] {
			Onelocalextrememin(inslice, q, r, outslice)
		}
	}
	return outslice
}

func Onelocalextrememax(inslice []float64, p int, r int, outslice []Localextreme) []Localextreme {
	q := int(math.Floor(float64((p + r) / 2)))
	if p <= r && q != r-1 && q > 1 {
		if inslice[q] >= inslice[q-1] && inslice[q] >= inslice[q+1] {
			locex := Localextreme{q, inslice[q], false, true}
			outslice = append(outslice, locex)
		}
		if inslice[q-1] <= inslice[q] && inslice[q] <= inslice[q+1] {
			Onelocalextrememax(inslice, q, r, outslice)
		}
		if inslice[q+1] <= inslice[q] && inslice[q] <= inslice[q-1] {
			Onelocalextrememax(inslice, p, q, outslice)
		}
		if inslice[q-1] >= inslice[q] && inslice[q+1] >= inslice[q] {
			Onelocalextrememax(inslice, q, r, outslice)
		}
	}
	return outslice
}

// Benchmarks and tests

func BenchmarkOnelocalextreme(b *testing.B) {
	ta := []float64{9.0, 7.0, 7.0, 2.0, 1.0, 2.0, 7.0, 5.0, 4.0, 7.0, 3.0, 4.0, 4.0, 8.0, 6.0, 9.0}
	for i := 0; i < b.N; i++ {
		Onelocalextreme(ta)
	}
}
