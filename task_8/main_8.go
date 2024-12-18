package main

type mutex struct {
	ch chan struct{}
}

func createMutex() *mutex {
	ch := make(chan struct{}, 1)
	return &mutex{
		ch: ch,
	}
}

func (mu *mutex) lock() {
	mu.ch <- struct{}{}
}

func (mu *mutex) unlock() {
	<-mu.ch
}
