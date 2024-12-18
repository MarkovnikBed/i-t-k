package main

import "testing"

func TestMerger(t *testing.T) {

	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()

	merged := merger(ch)

	if num := <-merged; num != 1 {
		t.Logf("ожидается получить 1 в результирующем канале, получено %d", num)
		t.Fail()
	}
}
