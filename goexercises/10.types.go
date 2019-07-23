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
	goexercises := Course{ Name: "Go", Slug: "go", Skills: [] string { "1", "2" } }
	fmt.Println(goexercises)

	goexercises1 := new(Course)
	goexercises1.Name = "Go1"
	goexercises1.Slug = "go1"
	goexercises1.Skills = [] string { "backend1" }
}