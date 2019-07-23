package main

import (
	"fmt"
)

func main() {
	getArray()
	getSlice()
}

func getArray() {
	var arr1 [2]string
	arr1[0] = "value1"
	arr1[1] = "value2"
	// arr1[2] = "value2" // This will throw exception
	fmt.Println(arr1)
}

func getSlice() {
	var slice1 []string
	slice1 = append(slice1, "mi", "slice", "1")
	fmt.Println(slice1)
}