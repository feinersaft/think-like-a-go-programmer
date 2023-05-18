package main

import (
	"fmt"
	"math"
)

func main() {
	raute := "#"
	maxRaut := 12 
	fmt.Println()
	for i := 0; i < maxRaut*2; i++ {
		for j := 0; j < (maxRaut - int(math.Abs(float64(maxRaut) - float64(i)))); j++ {
			fmt.Print(raute)
		}
		fmt.Print("\n")
	}
}
