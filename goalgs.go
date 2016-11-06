package main

import (
	"fmt"

	"github.com/paulidealiste/goalgs/rangen"
	"github.com/paulidealiste/goalgs/sortgen"
)

func main() {
	fmt.Println("Hello, goalgs!")
	tsl := rangen.Gorands(10, false, 10)
	fmt.Println(tsl)
	tsl = sortgen.Insertsort(tsl)
	fmt.Println(tsl)
}
