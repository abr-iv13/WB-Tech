// 13: Чем завершится данная программа?
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

/*
Данная программа выводит:

4
0
1
2
3
fatal error: all goroutines are asleep - deadlock!

Причина:
	Нужно передать указатель на объект WaitGroup, а не объект WaitGroup.
	Когда передается фактический WaitGroup, Go создает копию значения и вызывает Done() для копии.

Решение:
	Для решения дпнной броблемы нужно передать указатель, и каждая функция будет ссылаться на один и тот же WaitGroup.

	func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
	}
*/
