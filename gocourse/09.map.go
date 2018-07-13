package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMap())
	fmt.Println(getMap2())
	fmt.Println(getMap3("Sergio"))
}

func getMap()  map[string] int {
	mapTest := make(map[string] int)
	mapTest["key1"] = 1
	mapTest["key2"] = 100

	return mapTest
}

func getMap2()  map[string] int {
	mapTest := map[string] int {
		"Juan": 18,
		"Sergio": 24,
		"Julian": 30,
	}

	delete(mapTest, "Juan")

	return mapTest
}

func getMap3(name string)  int {
	mapTest := map[string] int {
		"Juan": 18,
		"Sergio": 24,
		"Julian": 30,
	}

	delete(mapTest, "Juan")

	return mapTest[name]
}