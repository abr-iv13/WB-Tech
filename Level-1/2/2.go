// 2: Написать программу, которая конкурентно рассчитает значение квадратов значений
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package main

import "fmt"

func main() {
	// arr := [5]int{2,4,6,8,10}
	arr := []int{2, 4, 6, 8, 10}

	// Создание канала
	ch := make(chan int)

	// Ф-ция рассчитывает значения кватров взятых из слайса arr
	go func(arr []int, ch chan int) {
		for _, v := range arr {
			ch <- v * v
		}
		close(ch)
	}(arr, ch)

	for v := range ch {
		fmt.Println(v)
	}
}
