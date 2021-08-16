//Конкретная стартегия
package eviction

import "fmt"

type Lru struct {
}

func (l *Lru) Evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}
