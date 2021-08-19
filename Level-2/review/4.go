// 4. Что выведет программа? Объяснить вывод программы.
package main

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)           //Решение: Закрыть канал,
		for i := 0; i < 10; i++ { //так как чтение из пустого канала блокирует до записи в этот канал или до его закрытия.
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}

/*
Вывод:
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!
*/
