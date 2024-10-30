package main

import "fmt"

func encrypt(words string) string {
    runes := []rune(words)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
   
    replacements := map[rune]rune{
        'a': '0',
        'e': '1',
        'i': '2',
        'o': '2',
        'u': '3',
    }
    for i, r := range runes {
        if newChar, exists := replacements[r]; exists {
            runes[i] = newChar
        }
    }
    
    
    encrypted := string(runes) + "aca"
    
    return encrypted
}

func main() {
    result := encrypt("apple")
    fmt.Println(result) 
}
