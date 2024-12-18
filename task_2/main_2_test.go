package main

import (
	"testing"
)

func TestGetRandSlice(t *testing.T) {
	if slice := getRandSlice(10); len(slice) != 10 {
		t.Logf("возвращён слайс неправильной длины")
		t.Fail()
	}
}

func TestSliceExample(t *testing.T) {
	if !compareSlice(sliceExample([]int{1, 1, 1, 2, 2}), []int{2, 2}) {
		t.Logf("sliceExample не возвращает ожидаемый слайс")
		t.Fail()
	}

	if !compareSlice(sliceExample([]int{1, 1, 1, 1, 1}), nil) {
		t.Logf("sliceExample не возвращает ожидаемый слайс")
		t.Fail()
	}
}

func TestAddElements(t *testing.T) {
	slice := []int{1, 2, 3}
	newSlice := addElements(slice, 4)
	if !compareSlice([]int{1, 2, 3, 4}, newSlice) {
		t.Logf("addElements не возвращает ожидаемый слайс по условию")
		t.Fail()
	}
}

func TestRemoveElements(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	newSlice := removeElement(slice, 0)
	if !compareSlice(newSlice, []int{2, 3, 4, 5, 6}) {
		t.Logf("removeElements не возвращает ожидаемый слайс по условию")
		t.Fail()
	}
}

func compareSlice(slice1 []int, slice2 []int) bool {

	if slice1 == nil && slice2 == nil {
		return true
	}

	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
