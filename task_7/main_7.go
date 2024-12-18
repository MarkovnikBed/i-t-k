package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func merger(chs ...chan int) chan int {
	merged := make(chan int)
	for _, ch := range chs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range ch {
				merged <- i
			}
		}()
	}
	return merged
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	go func() {
		defer close(ch1)
		ch1 <- 1
	}()

	go func() {
		defer close(ch2)
		ch2 <- 2
	}()
	go func() {
		defer close(ch3)
		ch3 <- 3
	}()
	merged := merger(ch1, ch2, ch3)
	go func() {
		wg.Wait()
		close(merged)
	}()

	for num := range merged {
		fmt.Println(num)
	}
}
