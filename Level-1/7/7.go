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
	mx := sync.Mutex{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			mx.Lock()
			defer mx.Unlock()
			exampleMap[i] = i
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i, v := range exampleMap {
		fmt.Println(i, v)
	}
}
