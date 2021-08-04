//8: Дана переменная int64. Написать программу которая устанавливает i-й бит в 1 или 0.
package main

import (
	"fmt"
	"strconv"
)

// Sets the bit at pos in the integer n.
func setBit(n int64, pos uint) int64 {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n.
func clearBit(n int64, pos uint) int64 {
	mask := ^(1 << pos)
	n &= int64(mask)
	return n
}

func main() {
	var i int64 = 123456
	var pos uint = 4

	fmt.Println(strconv.FormatInt(int64(i), 2))

	b := setBit(i, pos)
	fmt.Println(strconv.FormatInt(int64(b), 2))

	c := clearBit(i, 9)
	fmt.Println(strconv.FormatInt(int64(c), 2))

}
