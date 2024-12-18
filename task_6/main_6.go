package main

import (
	"fmt"
	"sync"
	"time"
)

type Generator struct {
	seed       int
	prevNumber int
}

var mu sync.Mutex
var wg sync.WaitGroup

func (g *Generator) getRandNum(maxSize int) int {
	mu.Lock()
	defer mu.Unlock()
	if maxSize == 0 {
		return 0
	}
	g.prevNumber = (g.prevNumber*5 + g.seed) % maxSize
	return g.prevNumber
}

func (g *Generator) SetSeed(seed int) {
	g.seed = seed
}

func NewGenerator() *Generator {
	return &Generator{}
}

func main() {
	var ch = make(chan int)
	generator := NewGenerator()
	generator.SetSeed(int(time.Now().Unix()))
	maxSizeOfRandInt := 1000
	for i := 0; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- generator.getRandNum(maxSizeOfRandInt)
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	for num := range ch {
		fmt.Print(num, " ")
	}
}
