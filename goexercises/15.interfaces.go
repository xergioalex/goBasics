package main

import "fmt"


type Course struct {
	Name string
	Slug string
	Skills [] string
}


func (p Course) Subscribe (name string) {
	fmt.Printf("The person %s has subscribed to %s course\n", name, p.Name)
}


type Carrer struct {
	Course
}

func (p Carrer) Subscribe (name string) {
	fmt.Printf("The person %s has subscribed to %s carrer\n", name, p.Name)
}

type Organization interface {
	Subscribe(name string)
}

func callSubscribe(p Organization) {
	p.Subscribe("XergioAleX")
}

func InterfaceTest() {
	myCourse := Course{ Name: "Go", Slug: "go", Skills: []string{"1", "2"} }
	// myCourse.Subscribe("XergioAleX")
	myCareer := new(Carrer)
	myCareer.Name = "Go 2"
	myCareer.Slug = "go2"
	myCareer.Skills = []string{"3", "4"}
	callSubscribe(myCourse)
	callSubscribe(myCareer)
}

func main() {
	InterfaceTest()
}