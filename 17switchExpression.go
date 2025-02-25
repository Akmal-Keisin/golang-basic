package main

import "fmt"

func main() {
	name := "Akmal"

	switch name {
	case "Akmal":
		fmt.Println("Halo Akmal")
	case "Keisin":
		fmt.Println("Halo Keisin")
	case "Alfateh":
		fmt.Println("Halo Alfateh")
	default:
		fmt.Println("Hello")
	}

	// Using short statement
	switch nameTest := "Alfateh"; nameTest {
	case "Akmal":
		fmt.Println("Halo Akmal")
	case "Keisin":
		fmt.Println("Halo Keisin")
	case "Alfateh":
		fmt.Println("Halo Alfateh")
	default:
		fmt.Println("Hello")
	}

	// Using condition
	score := 12
	switch {
	case score > 20:
		fmt.Println("Score is above 20")
	default:
		fmt.Println("Score is bellow 20")

	}
}
