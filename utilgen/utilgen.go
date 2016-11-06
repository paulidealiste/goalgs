//utilgen holds various utility function primarily for documenting the algs output and performance
package utilgen

import (
	"fmt"
	"time"
)

//Simple timetracker function called with defer at the onset of the function
func Timetracker(start time.Time, fname string) {
	elapsed := time.Since(start)
	fmt.Printf("Function %s ran for %s\n", fname, elapsed)
}
