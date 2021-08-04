//34: Написать программу, которая проверяет, что все символы в строке уникальные.
package main

import "fmt"

func unique(s string) bool {
	um := make(map[rune]bool)

	for _, i := range s {
		_, ok := um[i]
		if ok {
			return false
		}
		um[i] = true
	}
	return true
}

func main() {
	exampleOne := "hello"
	exampleTwo := "helo"

	fmt.Println(unique(exampleOne))
	fmt.Println(unique(exampleTwo))
}
