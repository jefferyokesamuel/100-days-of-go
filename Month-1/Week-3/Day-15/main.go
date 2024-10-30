package main

import (
    "fmt"
    "sort"

)

func sortString(s string) string {
    chars := []rune(s)
    sort.Slice(chars, func(i, j int) bool {
        return chars[i] < chars[j]
    })
    return string(chars)
}

func anagram(a string, b string) bool {

    if len(a) != len(b) {
        return false
    }
 
    sortedA := sortString(a)
    sortedB := sortString(b)
    
    return sortedA == sortedB
}

func main() {
    anagramTest := anagram("bat", "tab")
    fmt.Println(anagramTest) 
}
