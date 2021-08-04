//30: Удалить i-ый элемент из слайса.
package main

import (
	"fmt"
)

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func main() {
	s := []string{"A", "B", "C", "D", "E"}
	i := 1

	fmt.Println(remove(s, i))
}
