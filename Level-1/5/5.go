// 5: Написать программу, которая будет последовательно писать значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершиться.

package main

import (
	"fmt"
	"time"
)

var timeStop int = 4

// Ф-ция читает данные с канала и выводит stdout
func printChannelValue(ch <-chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}

//Ф-ция передает данные в канал c задежкой в одну секунду
func sendChannelValue(ch chan<- int) {
	list := 0
	for {
		ch <- list
		list++
		time.Sleep(time.Second)
	}

}

func main() {
	ch := make(chan int)

	go sendChannelValue(ch)
	go printChannelValue(ch)

	time.Sleep(time.Duration(timeStop) * time.Second)
	close(ch)
}
