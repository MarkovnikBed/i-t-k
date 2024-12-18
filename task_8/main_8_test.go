package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	var wg sync.WaitGroup
	var sum = 0

	mu := createMutex()

	increment := func() {
		defer wg.Done()
		mu.lock()
		sum++
		mu.unlock()
	}

	for range 10000 {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	if sum != 10000 {
		t.Logf("ожидалось 10000, получено %d", sum)
		t.Fail()
	}

	fmt.Println(sum)
}
