package main

import (
	"fmt"
)

func progressDays(days []int) int{
	var progressCount int
	for i := range days{
		if i > 0 {
		if days[i] > days[i-1]{
			progressCount ++
		}
	}
	}
	return progressCount
}

func main() {
	count := progressDays([]int{10, 11, 12, 9, 10})
	fmt.Println(count)
}