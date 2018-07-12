package structs

import "fmt"

func GetArray() {
	var arr1 [2] string
	arr2 := [3] int { 1, 2, 3 }
	arr1[0] = "array"
	arr1[1] = "array2"
	fmt.Println(arr1, arr2)
}

func GetSlice() {
	var slice1 [] string
	slice1 = append(slice1, "mi", "slice", "1")
	fmt.Println(slice1)
}