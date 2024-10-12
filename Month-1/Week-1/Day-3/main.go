package main

import (
	"fmt"
	"strings"
)

func getDiet(meals []string) (int, int) {
	var meatCount int
	var veggieCount int
	for _, v := range meals {
		if strings.Contains(v, "x") {
			meatCount++
		} else {
			veggieCount++
		}
	}
	return veggieCount, meatCount
}

func main() {
	veggieCount, meatCount := getDiet([]string{
		"--oooo-ooo--",
		"--xxxxxxxx--",
		"--o---",
		"-o-----o---x--",
		"--o---o-----",
	})
	fmt.Printf("Veggie Count: %v, Meat Count: %v\n", veggieCount, meatCount)

}
