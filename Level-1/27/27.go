//27: Написать программу, которая переворачивает слова в строке (snow dog sun - sun dog snow).
package main

import (
	"fmt"
	"strings"
)

func reverseWord(s string) string {
	words := strings.Split(s, " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	str := strings.Join(words, " ")
	return str
}

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWord(str))
}
