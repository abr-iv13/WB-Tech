//6: Какие существуют способы остановить выполнения горутины?
//Написать примеры использования.

package main

import "sync"

func main() {
	var wg sync.WaitGroup

	//Add увеличивает счетчик на 1
	wg.Add(1)

	ch := make(chan int)

	go func() {
		for {
			foo, ok := <-ch
			//Проверка на присутствие значение в канале
			if !ok {
				println("done")
				//Done уменьшает счетчик WaitGroup на один.
				wg.Done()
				return
			}
			println(foo)
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	//Wait блокирует исполнение потока до тех пор пока счетчик WaitGroup не обнулится.
	wg.Wait()
}
