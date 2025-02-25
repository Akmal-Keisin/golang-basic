package main

import "fmt"

func main() {
	runApp()
	runApp2(true)

	// Will not shown if the runApp2 is true
	fmt.Println("You can see me now")

}

func logging() {
	fmt.Println("logging() function has been called")
}

func runApp() {

	// logging will be called no matter if the runApp error\
	// this function with defer will be called in the end of the runApp()
	defer logging()
	fmt.Println("Run application")

	// this function will not be called if the runApp is error
	// logging()
}

func endApp() {
	fmt.Println("End App")

	// used to catch the panic message and keep the program running
	message := recover()
	fmt.Println("panic is happen", message)
}

func runApp2(err bool) {
	defer endApp()

	if err {
		// Used to stop the program
		// Usually used when panic condition is happen
		// defer will be called even the panic is executed
		panic("Ups Error")
	}

	fmt.Println("Application running")
}
