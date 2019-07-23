package main

import "fmt"
import "errors"

func Sum(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
		case string:
			return 0, errors.New("The value A is a string")
	}

	switch b.(type) {
		case string:
			return 0, errors.New("The value A is a string")
	}

	return a.(int) + b.(int), nil
}

func main() {
	number, err := Sum(50, 50)

	if err != nil {
		panic(err)
	}

	fmt.Println(number)
}