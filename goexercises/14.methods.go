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

func main() {
	myCourse1 := Course{ Name: "Go", Slug: "go", Skills: []string{"1", "2"} }
	myCourse1.Subscribe("XergioAleX")

	fmt.Println(myCourse1)
}