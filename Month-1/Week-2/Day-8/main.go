package main

import "fmt"


func main() {
	phoneMap := make(map[int][]string)
	phoneMap[2] = []string{"a", "b", "c"}
	phoneMap[3] = []string{"d", "e", "f"}
	phoneMap[4] = []string{"g", "h", "i"}
	phoneMap[5] = []string{"j", "k", "l"}
	phoneMap[6] = []string{"m", "n", "o"}
	phoneMap[7] = []string{"p", "q", "r", "s"}
	phoneMap[8] = []string{"t", "u", "v"}
	phoneMap[9] = []string{"w", "x", "y", "z"}

	possibleCombinations := func(a int, b int) []string{
		var combination []string
		for _, letterA := range phoneMap[a] {
			for _, letterB := range phoneMap[b] {
				combination = append(combination, letterA+letterB)
			}
		} 

		return combination
	}

	result := possibleCombinations(2,9)
	fmt.Println(result)
}