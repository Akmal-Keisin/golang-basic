package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Car struct {
	Name string
}

func (car Car) GetName() string {
	return car.Name
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func main() {
	person := Person{"Keisin"}
	SayHello(person)

	car := Car{"Lamborgini"}
	SayHello(car)
}
