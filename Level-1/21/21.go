//21: Написать программу, которая в конкурентном виде читает элементы из массива в stdout.
package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//Определяем группу в виде переменной wg sync.WaitGroup.
	wg := &sync.WaitGroup{}

	for i := 0; i < len(arr); i++ {
		//Вызов метода wg.Add() увеливает внутренний счетчик акетивных элементов на 1
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			//Вызов метода wg.Done() уменьшает внутренний счетчик активных элементов на единицу.
			defer wg.Done()
			//Вывести в stdOut
			fmt.Println(arr[i])
		}(i, wg)
	}
	//Вызыв метода Wait(), ожидает завершения всех горутин из группы wg.
	wg.Wait()
}
