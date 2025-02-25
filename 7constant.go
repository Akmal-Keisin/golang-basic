package main

import "fmt"

func main() {

	// Doesn't error even the constant is not used yet
	const firstName string = "Akmal"
	const lastName = "Keisin"

	// Error, cannot change the constant value
	// firstName = 'Edited'

	fmt.Println(firstName)

	// Declare multiple constant at once
	const (
		firstName1 string = "Martabak"
		lastName2         = "Kue Bandung"
	)
}
