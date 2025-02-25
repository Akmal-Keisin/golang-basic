package main

import "fmt"

func main() {

	// Logical Opeartion
	// && 	Logical "and".	Returns true if both statements are true	x < 5 &&  x < 10
	// || 	Logical "or".	Returns true if one of the statements is true	x < 5 || x < 4
	// !		Logical "not".	Reverse the result, returns false if the result is true	!(x < 5 && x < 10)

	var value1 bool = true
	var value2 bool = false

	var result1 = value1 && value2
	var result2 = value1 || value2
	var result3 = !value1

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
