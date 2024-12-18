package main

import (
	"testing"
)

func TestIntersect(t *testing.T) {
	testArrTrue := [][][]int{
		{{1, 2, 3, 4, 5, 6, 7, 8}, {4, 5, 6, 7, 8}, {4, 5, 6, 7, 8}},
		{{123, 45, 76, 56, 32, 223}, {45, 12, 13, 890, 33}, {45}},
		{{0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
	}
	testArrFalse := [][][]int{
		{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 10}, {}},
		{{100, 102, 23}, {110, 101, 22}, {}},
		{{2, 2, 2, 2, 2, 2, 2, 2}, {1, 1, 1, 1, 1, 1, 1, 1}, {}},
	}

	for _, v := range testArrTrue {
		ok, resSlice := intersect(v[0], v[1])
		if !ok {
			t.Logf("ожидается true, но получено false")
			t.Fail()
		}
		if !compareSlices(resSlice, v[2]) {
			t.Logf("результирующий слайс не соответсвует пересечению")
			t.Fail()
		}
	}

	for _, v := range testArrFalse {
		ok, resSlice := intersect(v[0], v[1])
		if ok {
			t.Logf("ожидается false, но получено true")
			t.Fail()
		}
		if !compareSlices(resSlice, v[2]) {
			t.Logf("результирующий слайс не должен содержать элементов")
			t.Fail()
		}
	}

}

func compareSlices(slice1 []int, slice2 []int) bool {
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
