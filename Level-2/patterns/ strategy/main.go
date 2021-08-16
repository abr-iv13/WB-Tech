package main

import (
	"github.com/abr-iv13/WB-Tech/tree/master/Level-2/patterns/strategy/eviction"
)

func main() {
	lfu := &eviction.Lfu{}
	cache := eviction.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &eviction.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &eviction.Fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")

}
