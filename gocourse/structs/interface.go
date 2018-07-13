package structs

// An interface for Course and Carrer
type Organization interface {
	Subscribe(name string)
}

func InterfaceTest() {
	goCourse := Course{ Name: "Go", Slug: "go", Skills: [] string { "1", "2" } }
	goCarrer := new(Carrer)
	goCarrer.Name = "GoCarrer"
	goCarrer.Slug = "go"
	callSubscribe(goCourse)
	callSubscribe(goCarrer)
}

func callSubscribe(p Organization) {
	p.Subscribe("XergioAleX")
}

