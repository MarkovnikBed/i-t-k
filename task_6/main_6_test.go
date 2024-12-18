package main

import (
	"testing"
	"time"
)

func TestGenerator(t *testing.T) {
	generator := NewGenerator()
	generator.SetSeed(int(time.Now().Unix()))
	var maxSizeOfRandInt int = 1
	for i := 0; i < 10; i++ {
		maxSizeOfRandInt *= 10
		if randNum := generator.getRandNum(maxSizeOfRandInt); randNum > maxSizeOfRandInt {
			t.Log("сгенерированное число превышает установленное ограничение")
			t.Fail()
		} else if randNum < 0 {
			t.Log("сгенерированное число не может быть отрицательным")
			t.Fail()
		}
	}

	if num := generator.getRandNum(0); num != 0 {
		t.Logf("ожидается 0, получено %v", num)
		t.Fail()
	}
}
