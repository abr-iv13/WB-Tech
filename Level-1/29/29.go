//29: Написать программу, которая перемножает, делит, складывает,
//вычитает 2 числовых переменных a,b, значение которые > 2^20
package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("24000000000000000000", 10)
	b.SetString("23000000000000000000", 10)

	result := new(big.Int)

	fmt.Println("Умножение :", result.Mul(a, b))
	fmt.Println("Деление  :", result.Div(a, b))
	fmt.Println("Cложение :", result.Add(a, b))
	fmt.Println("Вычитание:", result.Sub(a, b))
}
