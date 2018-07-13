package main

import (
	"fmt"
	"goBasics/gocourse/mypackage"
	"goBasics/gocourse/numbers"
	"goBasics/gocourse/structs"
	. "goBasics/gocourse/flow"
	"goBasics/gocourse/stringsc"
)

func main() {
	firstName := mypackage.GetName()
	fmt.Println("Your name is: ", firstName)

	stringUtf8 := mypackage.GetUnicode()
	fmt.Println("Unicode string: ", stringUtf8)

	sumResult := numbers.Sum(50, 50)
	fmt.Println("Sum result: ", sumResult)

	a, b, c := numbers.GetVariables()
	fmt.Println("Numbers a, b, c:", a, b, c)

	f32, f34 := numbers.GetFloat()
	fmt.Println("Numbers f32, f34:", f32, f34)

	structs.GetArray()
	structs.GetSlice()

	IfTest();
	SwitchTest();

	stringsc.Strings2()

	fmt.Println(structs.GetMap())
	fmt.Println(structs.GetMap2())
	fmt.Println(structs.GetMap3("Sergio"))

	structs.InterfaceTest()
}