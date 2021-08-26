//7: Реализовать конкурентную запись в map.
package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 100
	exampleMap := map[int]int{}
	wg := sync.WaitGroup{}

	//Мьютексы позволяют разграничить доступ к некоторым общим ресурсам, гарантируя,
	//что только одна горутина имеет к ним доступ в определенный момент времени.
	//И пока одна горутина не освободит общий ресурс, другая горутина не может с ним работать.
	mx := sync.Mutex{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			//Блокировка горутитины
			mx.Lock()
			//Разблокировка до завершения ф-ции
			defer mx.Unlock()
			//Запись данных в мапу
			exampleMap[i] = i
			//Вызов метода wg.Done() уменьшает внутренний счетчик активных элементов на единицу.
			wg.Done()
		}(i)
	}
	//Ожидание завершения всех горутин из группы wg:
	wg.Wait()

	//Чтение данных из мапы.
	for i, v := range exampleMap {
		fmt.Println(i, v)
	}
}
