package main

import "fmt"

func main() {
	raute := "#"
	n := 5
	fmt.Println()
	for i := 0; i < 5; i++ {
		for j := n; j > 0; j-- {
			fmt.Print(raute)
		}
		fmt.Print("\n")
		n -= 1
	}
}
