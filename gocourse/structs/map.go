package structs

func GetMap()  map[string] int {
	mapTest := make(map[string] int)
	mapTest["key1"] = 1
	mapTest["key2"] = 100

	return mapTest
}

func GetMap2()  map[string] int {
	mapTest := map[string] int {
		"Juan": 18,
		"Sergio": 24,
		"Julian": 30,
	}

	delete(mapTest, "Juan")

	return mapTest
}

func GetMap3(name string)  int {
	mapTest := map[string] int {
		"Juan": 18,
		"Sergio": 24,
		"Julian": 30,
	}

	delete(mapTest, "Juan")

	return mapTest[name]
}