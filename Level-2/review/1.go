//1. Что выведет программа? Объяснить вывод программы.
package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 78, 78, 79, 80} // array типа [5]int
	var b []int = a[1:4]            //создаем слайс []int со значениями [76 78 79]

	fmt.Println(b)
}

/*
Ответ: [77 78 79]
*/
