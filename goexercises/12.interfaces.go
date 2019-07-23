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


type Carrer struct {
	Course
}

func (p Carrer) Subscribe (name string) {
	fmt.Printf("La persona %s se ha registrado a la carrera %s\n", name, p.Name)
}

// An interface for Course and Carrer
type Organization interface {
	Subscribe(name string)
}

func callSubscribe(p Organization) {
	p.Subscribe("XergioAleX")
}


func main() {
	goexercises := Course{ Name: "Go", Slug: "go", Skills: [] string { "1", "2" } }
	goCarrer := new(Carrer)
	goCarrer.Name = "GoCarrer"
	goCarrer.Slug = "go"
	callSubscribe(goexercises)
	callSubscribe(goCarrer)
}