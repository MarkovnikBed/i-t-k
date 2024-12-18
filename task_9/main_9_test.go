package main

import "testing"

func TestConverter(t *testing.T) {
	chOne := make(chan uint)
	chTwo := make(chan float64)
	go func() {

		chOne <- 1
		chOne <- 2
		chOne <- 3
		chOne <- 0
		close(chOne)
	}()
	converter(chOne, chTwo)
	go func() {
		wg.Wait()
		close(chTwo)
	}()
	if floatNum := <-chTwo; floatNum != 1 {
		t.Logf("ожидалось преобразование в 1, получено %.2f", floatNum)
		t.Fail()
	}
	if floatNum := <-chTwo; floatNum != 8 {
		t.Logf("ожидалось преобразование в 8, получено %.2f", floatNum)
		t.Fail()
	}
	if floatNum := <-chTwo; floatNum != 27 {
		t.Logf("ожидалось преобразование в 27, получено %.2f", floatNum)
		t.Fail()
	}
	if floatNum := <-chTwo; floatNum != 0 {
		t.Logf("ожидалось преобразование в 0, получено %.2f", floatNum)
		t.Fail()
	}

}
