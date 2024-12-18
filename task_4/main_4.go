package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	expWords := map[string]struct{}{}
	for _, v := range slice2 {
		expWords[v] = struct{}{}
	}
	newSlice := getNewSlice(slice1, expWords)
	fmt.Println(newSlice)
}

func getNewSlice(slice []string, expWords map[string]struct{}) []string {
	var newSlice []string
	for _, v := range slice {
		_, ok := expWords[v]
		if !ok {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
