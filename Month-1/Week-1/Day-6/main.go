package main

import "fmt"


func nextPrime(num int) int {

	isPrime := func(n int) bool {
		if n <= 1 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	for {
		if isPrime(num) {
			return num
		}
		num++
	}
}

func main() {
	num := 29
	fmt.Printf("Next prime after %d is %d\n", num, nextPrime(num+1)) 
}
