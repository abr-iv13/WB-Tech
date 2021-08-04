//4: Реализовать набор из N воркеров, которые читают из канала произвольные данные и выводят в stdout. Данные в канал пишутся из главного потока.
//Необходима возможность выбора кол-во воркеров при старте, а также способ завершения работы всех воркеров.
package main

import (
	"fmt"
	"sync"
)

const COUNT = 7

func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan int)

	wg.Add(COUNT)
	for i := 0; i < COUNT; i++ {
		go worker(ch, wg)
	}

	for _, v := range arr {
		ch <- v
	}

	close(ch)
	wg.Wait()
}
