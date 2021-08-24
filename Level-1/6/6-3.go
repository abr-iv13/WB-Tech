//Закрытие ф-ции с помощью контекста...
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)

LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("finish from context timeout")
			break LOOP
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
