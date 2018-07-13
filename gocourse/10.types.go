package main

import (
	"fmt"
)

type Course struct {
	Name string
	Slug string
	Skills [] string
}

func main() {
	XergioAleX := Course{ Name: "Go", Slug: "go", Skills: [] string { "1", "2" } }
	fmt.Println(XergioAleX)
}