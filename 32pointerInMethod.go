package main

type Man struct {
	name string
}

func (man *Man) maried() {
	man.name = "Mr. " + man.name
}

func main() {
	person := Man{"John"}
	person.maried()
	println(person.name)
}