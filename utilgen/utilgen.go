// utilgen holds various utility function primarily for documenting the algs output and performance.
package utilgen

import (
	"errors"
	"fmt"
	"time"
)

// Simple timetracker function called with defer at the onset of the function.
func Timetracker(start time.Time, fname string) {
	elapsed := time.Since(start)
	fmt.Printf("Function %s ran for %s\n", fname, elapsed)
}

// Swap items in the supplied slice/tuple which sould be a pair of values.
func Swapitems(intuple []float64) ([]float64, error) {
	if len(intuple) != 2 {
		err := errors.New("Tuple (slice of length 2) is required for swapping.")
		panic(err)
	}
	intuple[0], intuple[1] = intuple[1], intuple[0]
	return intuple, nil
}
