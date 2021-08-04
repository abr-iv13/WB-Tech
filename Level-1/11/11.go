//11: Написать пересечение двух неупорядоченных массивов.
package main

import "fmt"

func intersection(s1, s2 []int) []int {
	inter := []int{}
	hash := make(map[int]bool)

	for _, e := range s1 {
		hash[e] = true
	}

	for _, e := range s2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}

	return inter
}

func main() {

	s1 := []int{6, 4, 8, 9, 3}
	s2 := []int{7, 6, 1, 8, 5, 3, 6, 9, 8}

	result := intersection(s1, s2)
	fmt.Println(result)

}
