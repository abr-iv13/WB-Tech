package main

func main() {
	lfu := &eviction.Lfu{}
	cache := eviction.initCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &eviction.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &eviction.fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")

}
