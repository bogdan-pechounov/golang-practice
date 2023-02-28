package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main1() {
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func(ch <-chan int) {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}(ch)
		go func(ch chan<- int) {
			ch <- 42
			wg.Done()
		}(ch)
	}
	wg.Wait()
}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		ch <- 34
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
