//Конкретная стратегия
package eviction

import "fmt"

type Lfu struct {
}

func (l *Lfu) Evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}
