package main

import (
	"fmt"
)

type Course struct {
	Name string
	Slug string
	Skills [] string
}

func (p Course) Subscribe (name string) {
	fmt.Printf("La persona %s se ha registrado al curso %s\n", name, p.Name)
}


func main() {
	goCourse := Course{ Name: "Go", Slug: "go", Skills: [] string { "1", "2" } }
	goCourse.Subscribe("Sergio")
}