package main

import "fmt"

func freeInmates(cells []int) int {
	if cells[0] == 0 {

		return 0
	}
	
	var freedInmates int
	for i := 0; i < len(cells); i++ {
		if cells[i] == 0 {
			continue
		}
		
		freedInmates++
		cells[i] = 0 // 


		for j := 0; j < len(cells); j++ {
			if cells[j] == 1 {
				cells[j] = 0
			} else {
				cells[j] = 1
			}
		}

		fmt.Println(cells) 
	}

	return freedInmates
}

func main() {
	cellsArray := []int{1, 1, 0, 0, 0, 1, 0}
	freedCriminals := freeInmates(cellsArray)
	fmt.Println(freedCriminals)
}