package main

import "testing"

func TestGetNewSlice(t *testing.T) {
	slice := []string{"1", "2", "4", "4", "5", "6", "goida"}
	expWords := map[string]struct{}{
		"2": {},
		"4": {},
	}
	newSlice := getNewSlice(slice, expWords)
	expSlice := []string{"1", "5", "6", "goida"}
	if !compareSlices(newSlice, expSlice) {
		t.Log("итоговый слайс не соответсвует условиям исключения элементов")
		t.Fail()
	}
}

func compareSlices(slice1 []string, slice2 []string) bool {
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
