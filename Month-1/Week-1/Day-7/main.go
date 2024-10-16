package main

import(
	"fmt"
	"sort"
)

func mergeArrays(a []int, b []int) []int {
	merged := append(a, b...)
	sort.Ints(merged)
	return merged
}

func main() {
	fmt.Println(mergeArrays([]int{1,2,3,4}, []int{1,6,7,8}))
}