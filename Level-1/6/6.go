//6: Какие существуют способы остановить выполнения горутины?
//Написать примеры использования.
package main

import "fmt"

func main() {
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Закрытие канала")
				return
			default:
			}
		}
	}()

	//Выход из горутины
	quit <- true
	fmt.Println("Горутина завершена")

}
