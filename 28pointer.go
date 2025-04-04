package main

import "fmt"

type SportCar struct {
	name       string
	types      string
	horsePower int
}

func main() {
	lamborgini := SportCar{
		name:       "Lamborgini",
		types:      "Sport",
		horsePower: 1000,
	}
	fmt.Println(lamborgini.name)
	fmt.Println(lamborgini.types)
	fmt.Println(lamborgini.horsePower)

	// Pointer to lamborgini
	var lamborgini2 *SportCar = &lamborgini
	lamborgini2.name = "Lamborgini 2"
	lamborgini2.types = "Sport 2"
	lamborgini2.horsePower = 2000

	fmt.Println("Lamborgini 2")
	fmt.Println(lamborgini2.name)
	fmt.Println(lamborgini2.types)
	fmt.Println(lamborgini2.horsePower)
	fmt.Println("Lamborgini 1")
	fmt.Println(lamborgini.name)
	fmt.Println(lamborgini.types)
	fmt.Println(lamborgini.horsePower)
}