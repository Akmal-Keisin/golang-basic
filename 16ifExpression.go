package main

import "fmt"

func main() {

	score := 10

	if score < 40 {
		fmt.Println("D")
	} else if score < 65 {
		fmt.Println("C")
	} else if score < 75 {
		fmt.Println("B")
	} else if score < 85 {
		fmt.Println("B+")
	} else if score < 95 {
		fmt.Println("A")
	} else {
		fmt.Println("A+")
	}

	if scoreTest := 10; scoreTest < 12 {
		fmt.Println("This is test if expreession using short statement")
	}
}
