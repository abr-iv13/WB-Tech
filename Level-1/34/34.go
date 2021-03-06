//34: Написать программу, которая проверяет, что все символы в строке уникальные.
package main

import "fmt"

func unique(s string) bool {
	um := make(map[rune]bool)

	for _, i := range s {
		//Проверка на наличии данного ключа в мапе
		_, ok := um[i]
		//Если ключ присутсвует возвращем false, иначе записываем ключ
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
