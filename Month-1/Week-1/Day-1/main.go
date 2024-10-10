package main

import (
	"fmt"
	"log"
)

func calcAge(age int) int{
	if age < 0{
		log.Fatal("Age can't be negative")
	}
	return 365 * age
}

func main() {
	fmt.Printf("Input Your Age:" )
	var age int
	fmt.Scan(&age)

	dayAge := calcAge(age)

	fmt.Printf("Your age in days is %v\n", dayAge)
}
