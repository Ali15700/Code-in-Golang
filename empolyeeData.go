package main

import (
	"fmt"
)

type Empolyee struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var emp Empolyee

	// empolyee specification
	fmt.Printf("Empolyee Name:")
	fmt.Scan(&emp.name)

	fmt.Printf("Empolyee Age:")
	fmt.Scan(&emp.age)

	fmt.Printf("Empolyee job:")
	fmt.Scan(&emp.job)

	fmt.Printf("Empolyee salary:")
	fmt.Scan(&emp.salary)
	fmt.Println("\n")
	printEmpolyee(emp)

}

func printEmpolyee(empo Empolyee) {
	fmt.Println("-----------Empolyee Details--------- ")
	fmt.Println("Name: ", empo.name)
	fmt.Println("Age: ", empo.age)
	fmt.Println("Job: ", empo.job)
	fmt.Println("Salary: ", empo.salary)
}
