package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/paulidealiste/goalgs/extremegen"
	"github.com/paulidealiste/goalgs/rangen"
	"github.com/paulidealiste/goalgs/sortgen"
)

func main() {
	color.Yellow("Hello, goalgs!")
	color.Yellow("--------------")
	// tsl := rangen.Gorands(10, false, 10)
	tsl := rangen.Boxmullerrands(20, 50, 20)
	color.Cyan("Original slice")
	fmt.Println(tsl)
	color.Green("goalgs sort algorithms")
	bsl := sortgen.Bubblesort(tsl)
	color.Cyan("Bubblesorted slice")
	fmt.Println(bsl)
	ssl := sortgen.Insertsort(tsl)
	color.Cyan("Insertsorted slice")
	fmt.Println(ssl)
	msl := sortgen.Mergesort(tsl, 0, len(tsl))
	color.Cyan("Mergesorted slice")
	fmt.Println(msl)
	color.Magenta("goalgs extremes algorithms")
	gss := extremegen.Findminmax(tsl)
	color.Cyan("Global extremes")
	fmt.Println(gss)
	fss := extremegen.Findlocalminmax(tsl)
	color.Cyan("Some of local extremes")
	fmt.Println(fss)
}
