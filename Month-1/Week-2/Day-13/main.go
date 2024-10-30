package main

import "fmt"

func calcWeight(items []int, bags int) bool{
	enoughBags := false

	bagCapacity := bags * 10

	itemSum := 0

	for _, v := range items {
		itemSum += v
	}

	if (itemSum > bagCapacity) {
		return enoughBags
	}

	enoughBags = true

	return enoughBags

}

func main() {
	groceryList := []int{1,2,3,4,5,5}

	shoppingSize := calcWeight(groceryList, 2) 

	fmt.Println(shoppingSize)
}