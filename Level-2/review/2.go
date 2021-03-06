//2. Что выведет программа? Объяснить вывод программы.
//Объяснить как работают defer’ы и их порядок вызовов.
package main

import (
	"fmt"
)

/*
Если отложенная функция является литералом функции(выражение, которое определяет неназванную функцию),
 а окружающая функция имеет именованные параметры результата,
 которые находятся в области видимости внутри литерала,
 отложенная функция может получить доступ к параметрам результата и изменить их,
 прежде чем они будут возвращены.*/
func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

/*
Если окружающая функция возвращается через явное return утверждение,
 отложенные функции выполняются после того,
 как этот return утверждение установит какие-либо параметры результата,
 но до того, как функция вернется к своему вызывающему.
*/
func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

/*
Ответ:
	2
	1

Оператор defer добавляет в стек вызов функции после ключевого слова defer

Очень важно учитывать, что после возвращения из внешней функции отло-
женные функции выполняются в порядке «последним пришел — первым вышел»
(Last In, First Out, LIFO).
*/
