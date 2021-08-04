//21: Написать программу, которая в конкурентном виде читает элементы из массива в stdout.
package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	wg := &sync.WaitGroup{}

	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			fmt.Println(arr[i])
			defer wg.Done()
		}(i, wg)
	}
	wg.Wait()
}
