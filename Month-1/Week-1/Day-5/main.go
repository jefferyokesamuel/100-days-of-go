package main

import (
	"strings"
)

func socksPairs(pairs string) int{
	socks := strings.Count(pairs, "A")
	return socks	
}

func main() {
	println(socksPairs("AA"))
}