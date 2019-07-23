package main

import "fmt"


type Course struct {
	Name string
	Slug string
	Skills [] string
}

func main() {
	myCourse1 := Course{ Name: "Go", Slug: "go", Skills: []string{"1", "2"} }
	myCourse2 := new(Course)
	myCourse2.Name = "Go2"
	myCourse2.Slug = "go2"
	myCourse2.Skills = []string{"3", "4"}
	fmt.Println(myCourse1)
	fmt.Println(myCourse2)
}