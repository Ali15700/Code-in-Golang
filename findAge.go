package main

import (
	"fmt"
)

func main() {

	var currentyear int = 2022
	var birthyear int = 0

	fmt.Print("Enter the year of birth:")
	fmt.Scan(&birthyear)

	age := currentyear - birthyear

	fmt.Println("Your age is:", age)
}
