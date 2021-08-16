//Интерфейс стратегии
package eviction

type EvictionAlgo interface {
	Evict(c *Cache)
}
