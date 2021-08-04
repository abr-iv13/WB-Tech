//32: Написать собственную функцию Sleep.
package main

import (
	"fmt"
	"time"
)

func sleep(s int) {
	<-time.After(time.Second * time.Duration(s))
}

func main() {
	fmt.Print("Hello")
	sleep(3)
	fmt.Print(" world \n")
}
