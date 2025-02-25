package main

import "fmt"

func main() {
	counter := 1
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	fmt.Println(counter)

	increment()
	increment()

	fmt.Println(counter)
}
