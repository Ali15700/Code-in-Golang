package main

import (
	"fmt"
)

func main() {

	day := ""

	fmt.Print("Enter the Name of Week Day:")
	fmt.Scan(&day)

	switch day {
	case "Monday":
		fmt.Print("Debug the code")
	case "Tuesday":
		fmt.Print("Create backend")
	case "Wednesday":
		fmt.Print("Create UI/UX ")
	case "Thursday":
		fmt.Print("Test the code")
	case "Friday":
		fmt.Print("Finalize the code")
	case "Saturday":
		fmt.Print("Holiday for today")
	case "Sunday":
		fmt.Print("Chill Day")

	default:
		fmt.Print("There's no week day of this name :(")
	}

}
