//33: Даны 2 канала - в первый пишутся рандомные числа после чего они проверяются на четность и отправляются во второй канал.
//Результаты работы из второго канала пишутся в stdout.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func check(in chan int, out chan int) {
	for el := range in {
		if el%2 == 0 {
			out <- el
		}
	}
	close(in)
}

func print(out chan int) {
	for el := range out {
		fmt.Println(el)
	}
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go check(in, out)
	go print(out)

	for i := 0; i < 10; i++ {
		in <- randInt(1, 100)
	}
}
