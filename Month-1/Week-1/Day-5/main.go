package main


func socksPairs(pairs string) int{
	counts := make(map[rune]int)

	for _, v := range pairs{
		if v >='A' && v <= 'Z'{
			counts[v]++
		}
	}
	
	pair := 0
	for _, count := range counts{
		pair += count / 2
	}
	return pair
}

func main() {
	println(socksPairs("AABBCC"))
}