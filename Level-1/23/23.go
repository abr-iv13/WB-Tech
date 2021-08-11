//23: Написать бинарный поиск встроенными методами языка.
package main

import (
	"fmt"
	"sort"
)

func checkBin(list []int, elem int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		//Вычислить номер среднего эл-та слайса
		mid := (low + high) / 2
		//Образер равен среднему элементу ?
		if list[mid] == elem {
			return mid
		}
		//Образец меньше среднего эл-та ?
		if list[mid] < elem {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(checkBin(arr, 3)) // 0
	// fmt.Println(checkBin([]int{1, 2, 3, 4, 5}, 4))  // 3
	// fmt.Println(checkBin([]int{1, 2, 3, 4, 5}, -1)) // -1

	fmt.Println(sort.SearchInts(arr, 3))
}
