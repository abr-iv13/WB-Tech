//9: Написать конвейер чисел. Даны 2 канала - в первый пишутся числа из массива,
// во второй пишется результат операции 2*x, после чего данные выводятся в
package main

func write(arr []int, in chan int) {
	for _, v := range arr {
		in <- v
	}
	close(in)
}

func calculate(in chan int, out chan int) {
	for v := range in {
		out <- v * 2
	}
	close(out)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	in := make(chan int)
	out := make(chan int)

	go write(arr, in)
	go calculate(in, out)

	for val := range out {
		println(val)
	}

}
