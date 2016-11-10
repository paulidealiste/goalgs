package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/paulidealiste/goalgs/rangen"
	"github.com/paulidealiste/goalgs/sortgen"
)

func main() {
	color.Yellow("Hello, goalgs!")
	color.Yellow("--------------")
	tsl := rangen.Gorands(10, false, 10)
	color.Cyan("Original slice")
	fmt.Println(tsl)
	color.Green("goalgs sort algorithms")
	ssl := sortgen.Insertsort(tsl)
	color.Cyan("Insertsorted slice")
	fmt.Println(ssl)
	msl := sortgen.Mergesort(tsl, 0, len(tsl))
	color.Cyan("Mergesorted slice")
	fmt.Println(msl)
}
