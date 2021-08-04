//19: К каким негативным последствиям может привести данный кусок кода и как это исправить?
package main

import "fmt"

var justString string

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {
		v += "A"
	}
	return v
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
	fmt.Println(justString)
}

/*
Возможно проблема в неграмотном распределение памяти, в функциб createHugeString аркументом получает 1024,
в ф-ции someFunc забираем v[:100] c 0 до 99.
*/
