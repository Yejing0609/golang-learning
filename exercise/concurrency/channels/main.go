package main

import (
	"sync"
	"time"
)

func main() {
	c := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		c <- `foo`
	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Second * 1)
		println(`message: ` + <-c)
	}()

	wg.Wait()
}
