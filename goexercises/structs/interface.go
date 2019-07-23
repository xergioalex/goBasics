package structs

// An interface for Course and Carrer
type Organization interface {
	Subscribe(name string)
}

func InterfaceTest() {
	goexercises := Course{ Name: "Go", Slug: "go", Skills: [] string { "1", "2" } }
	goCarrer := new(Carrer)
	goCarrer.Name = "GoCarrer"
	goCarrer.Slug = "go"
	callSubscribe(goexercises)
	callSubscribe(goCarrer)
}

func callSubscribe(p Organization) {
	p.Subscribe("XergioAleX")
}

