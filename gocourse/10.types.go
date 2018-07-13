package main

import (
	"fmt"
)

type Course struct {
	Name string
	Slug string
	Skills [] string
}
// Career contiene los datos de una carrera
type Career struct {
	Course
}


func main() {
	goCourse := Course{ Name: "Go", Slug: "go", Skills: [] string { "1", "2" } }
	fmt.Println(goCourse)

	goCourse1 := new(Course)
	goCourse1.Name = "Go1"
	goCourse1.Slug = "go1"
	goCourse1.Skills = [] string { "backend1" }
}