package main

import (
	"fmt"
)

func main() {
	var car = make(map[string]string)
	car["Car Owner"] = "Ali Raza"
	car["Car brand"] = "Bugatti"
	car["Car model"] = "Chiron"
	car["Make year"] = "2021"

	fmt.Printf(" Details %v\n", car)

}
