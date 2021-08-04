//14: Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func set(arr []string) (result []string) {
	setMap := make(map[string]bool)
	// result := []string{}

	for _, v := range arr {
		setMap[v] = true
	}

	for i := range setMap {
		result = append(result, i)
	}

	// return result
	return
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	result := set(arr)

	fmt.Println(result)
}
