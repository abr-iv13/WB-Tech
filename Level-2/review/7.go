//7. Что выведет программа? Объяснить вывод программы.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		//Итерация по входящим rest параметрам
		for _, v := range vs {
			//Передача значения итерированного парамета в канал
			c <- v
			//Рандомное ожидание
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		//Закрытие канала
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

/*
Операция приема по закрытому выполняется немедленно,
 получая нулевое значение типа элемента после получения любых ранее отправленных значений.
*/
