package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter something: ")
    input, _ := reader.ReadString('\n') 
	input = strings.TrimSpace(input)

	words := strings.Split(input, " ")
	for i, word := range words {
		if word == "Nemo"{
			fmt.Printf("I found Nemo at %v\n", i)
		}	
	}
}