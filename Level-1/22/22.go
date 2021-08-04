//22: Написать быструю сортировку встроенными методами языка.
package main

import "fmt"

//Ф-ция quicksort принимает слайс int
//Возвращает салайс int
func quicksort(list []int) []int {
	//Если размер слайса меньше двух, завершить программу
	if len(list) < 2 {
		return list
	} else {
		//Опорная переменная
		pivot := list[0]
		//Меньше
		var less = []int{}
		//Больше
		var greater = []int{}
		//Итерация входящего параметра типа слайс int
		for _, num := range list[1:] {
			if pivot > num {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}
		//Рекурсия
		less = append(quicksort(less), pivot)
		greater = quicksort(greater)
		return append(less, greater...)
	}
}

func main() {
	fmt.Println(quicksort([]int{-1, 9, 7, 8, 4, 5, 2, 3}))
}
