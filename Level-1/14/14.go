//14: Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func set(arr []string) (result []string) {
	setMap := make(map[string]bool)
	// result := []string{}

	//Итерация по входящему параметру тип []string
	for _, v := range arr {
		//заполняем мапу уникальными значениями
		setMap[v] = true
	}

	//Итерация мапы
	for k, _ := range setMap {
		//Все ключи аппендим в слайс result
		result = append(result, k)
	}

	// return result
	return
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(set(arr))
}
