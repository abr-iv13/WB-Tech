//9: Написать конвейер чисел. Даны 2 канала - в первый пишутся числа из массива,
//во второй пишется результат операции 2*x, после чего данные выводятся в stdout
package main

//Ф-ция write щаписывает в канал числа из входящего массива
func write(arr []int, in chan int) {
	//Закрытие канал
	defer close(in)
	for _, v := range arr {
		in <- v
	}
}

//Функция caluclate записывает в канал результаты операции 2 * x
func calculate(in chan int, out chan int) {
	//Закрытие канала
	defer close(out)

	for v := range in {
		out <- v * 2
	}
}

func main() {
	//Входящий массив
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//Создaние каналов на вход и на выход
	in := make(chan int)
	out := make(chan int)

	go write(arr, in)
	go calculate(in, out)

	//Чтение из канала калькулированные значение и вывод stdout
	for val := range out {
		println(val)
	}

}
