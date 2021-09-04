package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const COUNT = 7

func NewClient(t int) *http.Client {
	return &http.Client{
		Timeout: time.Second * time.Duration(t),
	}
}

func worker(ch chan int, wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()

	for v := range ch {
		url := fmt.Sprintf("https://webhook.site/ccbf76cd-5702-45dd-a650-23ae7fc6e7d2?id=%d", v)

		resp, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resp)

	}
}

func main() {
	inc := 0
	client := NewClient(10)
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	for i := 0; i < COUNT; i++ {
		wg.Add(1)
		go worker(ch, wg, client)
	}

	for i := 0; i < 10; i++ {
		inc++
		ch <- inc
	}

	close(ch)

	wg.Wait()
}
