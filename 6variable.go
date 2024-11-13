package main

import "fmt"

func main() {
	var hello string

	hello = "Hello World"
	fmt.Println(hello)

	hello = "Hello Go"
	fmt.Println(hello)

	var world = 123
	fmt.Println(world)

	// This is error, can't fill some variable with other data type
	// world = "hello"

	world = 321
	fmt.Println(world)

	thisIsAlsoVariable := "Hello Variable"
	fmt.Println(thisIsAlsoVariable)

	thisIsAlsoVariable = "Variable Changed"
	fmt.Println(thisIsAlsoVariable)

	// Multiple declaration variable
	var (
		firstVar  = "This is the first variable"
		secondVar = "This is the second variable"
	)

	fmt.Println(firstVar)
	fmt.Println(secondVar)

	// Cannot declare a variable without using it
	// var errorVariable = "This is error"
}
