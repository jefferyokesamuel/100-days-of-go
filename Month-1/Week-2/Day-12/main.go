package main

import "fmt"

func detectType(a []int) string{

	for i := range a{
		if i > 2{

		if a[i-2] < a[i-1] && a[i+1] > a[i+2] {
			return "mountain"
		}
		if a[i-2] > a[i-1] && a[i+1] < a[i+2] {
			return "valley"
		}
	}
		
	}

	return "Null"
}

func main() {
	nums := []int{10, 9, 8, 7, 2, 3, 4, 5}
	varType := detectType(nums)

	fmt.Println(varType)
}