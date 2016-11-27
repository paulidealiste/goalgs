// extremegen holds implementations of various extreme findining algorithms
package extremegen

import (
	"math"
	"time"

	"github.com/paulidealiste/goalgs/sortgen"
	"github.com/paulidealiste/goalgs/utilgen"
)

type Extreme struct {
	index    int
	value    float64
	min, max bool
}

type ExtremeDiff struct {
	index    []int
	value    float64
	values   []float64
	min, max bool
}

type ExtremeSlice struct {
	index  []int
	sum    float64
	target []float64
}

// Extremediff finds the minimal and maximal difference between any two consequtive members of the
// input array and returns the list of two struct element holding the relevant information
func Extremediff(inslice []float64) []ExtremeDiff {
	defer utilgen.Timetracker(time.Now(), "Leastdiff")
	var outslice []ExtremeDiff
	minDf := ExtremeDiff{nil, math.MaxFloat64, nil, true, false}
	maxDf := ExtremeDiff{nil, 0.0, nil, false, true}
	sortslice := sortgen.Insertsort(inslice)
	for i := 0; i < len(sortslice)-1; i++ {
		indiff := math.Abs(sortslice[i] - sortslice[i+1])
		if indiff > maxDf.value {
			maxDf.values = []float64{sortslice[i], sortslice[i+1]}
			maxDf.value = indiff
		}
		if indiff < minDf.value {
			minDf.values = []float64{sortslice[i], sortslice[i+1]}
			minDf.value = indiff
		}
	}
	minDf.index = utilgen.Retind(inslice, minDf.values)
	maxDf.index = utilgen.Retind(inslice, maxDf.values)
	outslice = append(outslice, minDf)
	outslice = append(outslice, maxDf)
	return outslice
}

// Maxminsubarray represents the algorithm for finding the nonempty subarray
// of a target array whose values have the largest/smallest sum.
func Maxminsubarray(inslice []float64) []ExtremeSlice {
	defer utilgen.Timetracker(time.Now(), "Maxminsubarray")
	outslice := maxminsub(inslice)
	return outslice
}

func maxminsub(inslice []float64) []ExtremeSlice {
	var outslice []ExtremeSlice
	mmchan := make(chan ExtremeSlice, 1)
	mnchan := make(chan ExtremeSlice, 1)
	if len(inslice) > 1 {
		go innersub(inslice, mmchan, false)
		go innersub(inslice, mnchan, true)
		outslice = append(outslice, <-mmchan)
		outslice = append(outslice, <-mnchan)
	} else {

	}
	return outslice
}

func innersub(inslice []float64, mmchan chan ExtremeSlice, ismin bool) {
	var innerext ExtremeSlice
	var (
		cs   int
		ce   int
		csum float64
	)
	for ce != len(inslice) {
		csum += inslice[ce]
		if ismin == true {
			if csum < innerext.sum {
				innerext.index = []int{cs, ce}
				innerext.target = inslice[cs:ce]
				innerext.sum = csum
			}
			if csum > 0 {
				cs = ce + 1
				csum = 0
			}
		} else {
			if csum > innerext.sum {
				innerext.index = []int{cs, ce}
				innerext.target = inslice[cs:ce]
				innerext.sum = csum
			}
			if csum <= 0 {
				cs = ce + 1
				csum = 0
			}
		}
		ce++
	}
	mmchan <- innerext
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
	outslice := make([]Extreme, 0)
	outslice = onelocalextrememin(inslice, 1, len(inslice)-1, outslice)
	outslice = onelocalextrememax(inslice, 1, len(inslice)-1, outslice)
	return outslice
}

func onelocalextrememin(inslice []float64, p int, r int, outslice []Extreme) []Extreme {
	if p < r {
		if inslice[p] <= inslice[p-1] && inslice[p] <= inslice[p+1] {
			locex := Extreme{p, inslice[p], true, false}
			outslice = append(outslice, locex)
		}
		return onelocalextrememin(inslice, p+1, r, outslice)
	}
	return outslice
}

func onelocalextrememax(inslice []float64, p int, r int, outslice []Extreme) []Extreme {
	if p < r {
		if inslice[p] >= inslice[p-1] && inslice[p] >= inslice[p+1] {
			locex := Extreme{p, inslice[p], true, false}
			outslice = append(outslice, locex)
		}
		return onelocalextrememax(inslice, p+1, r, outslice)
	}
	return outslice
}
