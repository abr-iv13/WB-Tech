//11: Написать пересечение двух неупорядоченных массивов.
package main

import "fmt"

func intersection(s1, s2 []int) []int {
	inter := []int{}              //Слайс пересечений
	tempMap := make(map[int]bool) //Временная мапа для сравнениния

	//Итерация слайса
	for _, v := range s1 {
		tempMap[v] = true //Присваиваем значение из слайса в ключ мап со значением true
	}

	for _, v := range s2 {
		//Проверка на присутсвие ключей у мапы из слайса s2
		//Если ключ присутсвует добавляем значение в слайс inter
		if tempMap[v] {
			inter = append(inter, v)
		}
	}

	return inter
}

func main() {
	s1 := []int{6, 4, 8, 9, 3}
	s2 := []int{7, 6, 1, 8, 5, 3, 6, 9, 8}

	fmt.Println(intersection(s1, s2))
}
