package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(intersect(a, b))
}

func intersect(a []int, b []int) (bool, []int) {

	var maxLenArr []int
	var minLenArr []int
	if len(a) >= len(b) {
		maxLenArr = a
		minLenArr = b
	} else {
		maxLenArr = b
		minLenArr = a
	}

	newMap := map[int]struct{}{}
	for _, v := range maxLenArr {
		newMap[v] = struct{}{}
	}

	var resSlice []int
	isInterSect := false

	for _, v := range minLenArr {
		if _, ok := newMap[v]; ok {
			resSlice = append(resSlice, v)
			isInterSect = true
		}
	}

	return isInterSect, resSlice
}
