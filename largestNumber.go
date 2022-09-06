package main

import (
	"fmt"
)

func main() {

	var array1 = [9]int{22, 142, 21, 342, 4325, 324, 532, 432, 234}
	largee := 0
	for i := 0; i < 8; i++ {
		if array1[i] > largee {

			largee = array1[i]
		}
	}
	fmt.Print("Largest Value:", largee)
}
