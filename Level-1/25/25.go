//25: Написать свою структуру счетчик, которая будет инкрементировать и выводить значения в конкурентной среде.
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	index int
	mu    sync.RWMutex
}

func (c *Counter) increment() {
	c.mu.Lock()
	c.index++
	c.mu.Unlock()
}

func (c *Counter) print() {
	c.mu.RLock()
	fmt.Println(c.index)
	c.mu.RUnlock()
}

func main() {
	c := new(Counter)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			c.increment()
			c.print()
			defer wg.Done()
		}()
	}
	wg.Wait()
}
