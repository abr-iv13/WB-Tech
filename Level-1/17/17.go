//17: Написать программу, которая в рантайме способна определить тип переменной
// — int, string, bool, channel из переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)

func checkType(param interface{}) {
	fmt.Println(reflect.TypeOf(param))

	//switch v := param.(type) {
	//case int:
	//	fmt.Println("int")
	//case string:
	//	fmt.Println("string")
	//case bool:
	//	fmt.Println("bool")
	//case chan int:
	//	fmt.Println("chan int")
	//default:
	//	fmt.Println("Переменная не определенна", v)
	//}
}

func main() {
	tChannel := make(chan int)
	tString := "Hello"
	tInt := 125
	tBool := true

	checkType(tChannel)
	checkType(tString)
	checkType(tInt)
	checkType(tBool)
}
