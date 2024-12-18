package main

import (
	"fmt"
	"math/rand"
)

func main() {
	originalSlice := getRandSlice(10)
	fmt.Println(originalSlice)
	fmt.Println(sliceExample(originalSlice))
	fmt.Println(addElements(originalSlice, 12))
	fmt.Println(removeElement(originalSlice, 5))
}

func getRandSlice(count int) []int {
	var slice []int
	for range count {
		slice = append(slice, rand.Intn(1000))
	}
	return slice
}

func sliceExample(originalSlice []int) []int {
	var slice []int
	for _, v := range originalSlice {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	return slice
}

func addElements(originalSlice []int, num int) []int {
	return append(originalSlice, num)
}

func removeElement(originalSlice []int, index int) []int {
	for i := range originalSlice {
		if i == index {
			return append(originalSlice[:i], originalSlice[i+1:]...)
		}
	}
	return originalSlice
}
