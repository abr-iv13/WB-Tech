package main

import (
	"fmt"
	"sync"
)

type CustomSyncMap struct {
	channel   chan struct{}
	customMap map[int]int
}

func (cm *CustomSyncMap) WriteMapValue(wg *sync.WaitGroup, val int) {
	defer wg.Done()
	cm.channel <- struct{}{}
	cm.customMap[val] = val
	<-cm.channel
}

func (cm *CustomSyncMap) ReadMapValue(wg *sync.WaitGroup) {
	defer wg.Done()
	cm.channel <- struct{}{}
	fmt.Println(cm.customMap)
	<-cm.channel
}

func main() {
	customSyncMap := &CustomSyncMap{
		channel:   make(chan struct{}, 1),
		customMap: make(map[int]int),
	}

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go customSyncMap.WriteMapValue(wg, i)
		go customSyncMap.ReadMapValue(wg)
	}

	wg.Wait()
}
