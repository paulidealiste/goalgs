package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/paulidealiste/goalgs/datagen"
	"github.com/paulidealiste/goalgs/rangen"
	"github.com/paulidealiste/goalgs/sortgen"
)

func main() {
	color.Yellow("Hello, goalgs!")
	color.Yellow("--------------")
	tsl := rangen.Boxmullerrands(20, 50, 100)
	color.Cyan("Original slice")
	fmt.Println(tsl)
	// color.Green("goalgs sort algorithms")
	// bsl := sortgen.Bubblesort(tsl)
	// color.Cyan("Bubblesorted slice")
	// fmt.Println(bsl)
	// ssl := sortgen.Insertsort(tsl)
	// color.Cyan("Insertsorted slice")
	// fmt.Println(ssl)
	// msl := sortgen.Mergesort(tsl, 0, len(tsl))
	// color.Cyan("Mergesorted slice")
	// fmt.Println(msl)
	qsl := sortgen.Heapsort(tsl)
	color.Cyan("Heapsorted slice")
	fmt.Println(qsl)
	// color.Blue("goalgs random permutation algorithms")
	// color.Cyan("Randomize by sort")
	// rpsl := sortgen.Sortpermute(tsl)
	// fmt.Println(rpsl)
	// color.Cyan("Randomize in-place")
	// ppsl := sortgen.Inplacepermute(tsl)
	// fmt.Println(ppsl)
	// color.Magenta("goalgs extremes algorithms")
	// gss := extremegen.Findminmax(tsl)
	// color.Cyan("Global extremes")
	// fmt.Println(gss)
	// fss := extremegen.Findlocalminmax(tsl)
	// color.Cyan("All of local extremes")
	// fmt.Println(fss)
	// ass := extremegen.Maxminsubarray(tsl)
	// color.Cyan("Maximum and minimum subarrays (if any)")
	// fmt.Println(ass)
	// sss := extremegen.Extremediff(tsl)
	// color.Cyan("Extreme differences between two elements of array")
	// fmt.Println(sss)
	color.Green("goalgs data structures")
	ssd := datagen.Heapgen(tsl)
	color.Cyan("Slice representation as max-heap (.inslice of the struct)")
	fmt.Println(ssd)
}
