// extremegen holds implementations of various extreme findining algorithms
package extremegen

import (
	"math"
	"testing"
	"time"

	"github.com/paulidealiste/goalgs/utilgen"
)

type Extreme struct {
	index    int
	value    float64
	min, max bool
}

// Maximumsubarray represents the algorithm for finding the nonempty subarray
// of a target array whose values have the largest sum.
func Maximumsubarray(inslice []float64) []float64 {

	return inslice
}

// Findminmax proceeds to find global extremes of the supplied array returning them as structs
// representing original index, the vaue of the extreme and the flag indicating the extreme type.
func Findminmax(inslice []float64) []Extreme {
	defer utilgen.Timetracker(time.Now(), "Findminmax")
	min := Extreme{0, 0, true, false}
	max := Extreme{0, 0, false, true}
	onefindmax(inslice, 0, len(inslice)-1, &max)
	onefindmin(inslice, 0, len(inslice)-1, &min)
	outslice := []Extreme{min, max}
	return outslice
}

func onefindmax(inslice []float64, p int, r int, max *Extreme) {
	if p == r {
		max.index = p
		max.value = inslice[p]
	} else {
		onefindmax(inslice, p+1, r, max)
		if inslice[p] > max.value {
			max.index = p
			max.value = inslice[p]
		} else {
			max = max
		}
	}
}

func onefindmin(inslice []float64, p int, r int, min *Extreme) {
	if p == r {
		min.index = p
		min.value = inslice[p]
	} else {
		onefindmin(inslice, p+1, r, min)
		if inslice[p] < min.value {
			min.index = p
			min.value = inslice[p]
		} else {
			min = min
		}
	}
}

// Findlocalminmax finds either/or a list of called for local extremes and returns the corresponding
// array of structs representing original index, the value of the detected local extreme and the flag
// indicating the type of the extreme.
func Findlocalminmax(inslice []float64) []Extreme {
	defer utilgen.Timetracker(time.Now(), "Findlocalminmax")
	var outslice []Extreme
	outslice = onelocalextrememin(inslice, 0, len(inslice)-1, outslice)
	outslice = onelocalextrememax(inslice, 0, len(inslice)-1, outslice)
	return outslice
}

func onelocalextrememin(inslice []float64, p int, r int, outslice []Extreme) []Extreme {
	q := int(math.Floor(float64((p + r) / 2)))
	if p <= r && q != r-1 && q > 1 {
		if inslice[q] <= inslice[q-1] && inslice[q] <= inslice[q+1] {
			locex := Extreme{q, inslice[q], true, false}
			outslice = append(outslice, locex)
			return onelocalextrememin(inslice, p, q, outslice)
		}
		if inslice[q-1] >= inslice[q] && inslice[q] >= inslice[q+1] {
			return onelocalextrememin(inslice, q, r, outslice)
		}
		if inslice[q+1] >= inslice[q] && inslice[q] >= inslice[q-1] {
			return onelocalextrememin(inslice, p, q, outslice)
		}
		if inslice[q-1] <= inslice[q] && inslice[q+1] <= inslice[q] {
			return onelocalextrememin(inslice, q, r, outslice)
		}
	}
	return outslice
}

func onelocalextrememax(inslice []float64, p int, r int, outslice []Extreme) []Extreme {
	q := int(math.Floor(float64((p + r) / 2)))
	if p <= r && q != r-1 && q > 1 {
		if inslice[q] >= inslice[q-1] && inslice[q] >= inslice[q+1] {
			locex := Extreme{q, inslice[q], false, true}
			outslice = append(outslice, locex)
			return onelocalextrememax(inslice, q, r, outslice)
		}
		if inslice[q-1] <= inslice[q] && inslice[q] <= inslice[q+1] {
			return onelocalextrememax(inslice, q, r, outslice)
		}
		if inslice[q+1] <= inslice[q] && inslice[q] <= inslice[q-1] {
			return onelocalextrememax(inslice, p, q, outslice)
		}
		if inslice[q-1] >= inslice[q] && inslice[q+1] >= inslice[q] {
			return onelocalextrememax(inslice, q, r, outslice)
		}
	}
	return outslice
}

// Benchmarks and tests

func BenchmarkFindlocalminmax(b *testing.B) {
	ta := []float64{9.0, 7.0, 7.0, 2.0, 1.0, 2.0, 7.0, 5.0, 4.0, 7.0, 3.0, 4.0, 4.0, 8.0, 6.0, 9.0}
	for i := 0; i < b.N; i++ {
		Findlocalminmax(ta)
	}
}
