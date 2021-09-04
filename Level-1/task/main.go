package main

import (
	"fmt"
	"sync"
)

type CustomSyncMap struct {
	channel   chan struct{}
	customMap map[int]int
}

func NewCustomSyncMap() *CustomSyncMap {
	return &CustomSyncMap{
		channel:   make(chan struct{}, 1),
		customMap: make(map[int]int),
	}
}

func (cm *CustomSyncMap) Read(wg *sync.WaitGroup) {
	defer wg.Done()
	cm.channel <- struct{}{}
	fmt.Println(cm.customMap)
	<-cm.channel
}

func (cm *CustomSyncMap) Write(wg *sync.WaitGroup, val int) {
	defer wg.Done()
	cm.channel <- struct{}{}
	cm.customMap[val] = val
	<-cm.channel
}

func main() {
	customSyncMap := NewCustomSyncMap()
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go customSyncMap.Write(wg, i)
		go customSyncMap.Read(wg)
	}

	wg.Wait()
}
