package main

import (
	"fmt"
	"math"
	"sync"
)

var wg sync.WaitGroup

func main() {
	chOne := make(chan uint)
	chTwo := make(chan float64)
	go func() {
		chOne <- 1
		chOne <- 2
		chOne <- 3
		chOne <- 4
		chOne <- 5
		chOne <- 6
		chOne <- 7
		chOne <- 8
		chOne <- 9
		chOne <- 10
		close(chOne)
	}()

	converter(chOne, chTwo)
	go func() {
		wg.Wait()
		close(chTwo)
	}()
	for floatNum := range chTwo {
		fmt.Println(floatNum)
	}
}

func converter(chOne <-chan uint, chtwo chan<- float64) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range chOne {

			chtwo <- math.Pow(float64(num), 3)

		}

	}()
}
