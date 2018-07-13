package main

import (
	"fmt"
)

type CustomCourse struct {
	Name string
	Slug string
	Skills [] string
}

// Course contiene los datos generales de un curso
type Course struct {
	Name        string
	Description string
	Slug        string
	Teacher Instructor
	Skills      []string
}

// Instructor contiene los datos de un instructor
type Instructor struct {
	FirstName     string
	LastName      string
	Twitter       string
	ProfilePicURL string
	Age           int
	Country       string
}

// Career contiene los datos de una carrera
type Career struct {
	Name        string
	Slug        string
	Description string
	Courses     []Course
}


func main() {
	goCourse := CustomCourse{ Name: "Go", Slug: "go", Skills: [] string { "1", "2" } }
	fmt.Println(goCourse)

	goCourse1 := new(CustomCourse)
	goCourse1.Name = "Go1"
	goCourse1.Slug = "go1"
	goCourse1.Skills = [] string { "backend1" }
}