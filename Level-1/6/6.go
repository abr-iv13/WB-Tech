//6: Какие существуют способы остановить выполнения горутины?
//Написать примеры использования.
package main

import "fmt"

func main() {
	//Создает канал с типом пустой структуры, для уведомления о закрытие struct{}{} весит 0 байт.
	quit := make(chan struct{})

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
	quit <- struct{}{}
	fmt.Println("Горутина завершена")

}
