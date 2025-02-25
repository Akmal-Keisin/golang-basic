package main

import "fmt"

type Filter func(string) string
type Blacklist func(string) bool

func main() {
	// Basic function
	greetings()

	// Function with parameter and return value
	fmt.Println(summary(10, 20))

	// Function with multiple return value
	firstName, secondName := greetings2("Akmal", "Keisin")
	fmt.Println(firstName)
	fmt.Println(secondName)

	// Ignoree function with multiple value
	firstName1, _ := greetings2("Akmal1", "Keisin1")
	fmt.Println(firstName1)

	// Function with named return value
	firstNameTest, middleNameTest, lastNameTest := greetings3()
	fmt.Println(firstNameTest)
	fmt.Println(middleNameTest)
	fmt.Println(lastNameTest)

	// Variadic function/variable argument
	fmt.Println(summary2(10, 20, 30, 40, 50))
	var numbers []int = []int{10, 20, 30, 40, 50}
	fmt.Println(summary2(numbers...))

	// Function as value
	sayHello := greetings
	sayHello()

	// Function as parameter
	sayHelloWithFilter("Akmal", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	// Function type declaration
	sayHelloWithFilter2("Akmal", spamFilter)
	sayHelloWithFilter2("Anjing", spamFilter)

	// Anonymous function
	registerUser("Akmal", func(name string) bool {
		if name == "Akmal" {
			return true
		} else {
			return false
		}
	})
}

func greetings() {
	fmt.Println("Hello")
}

func summary(firstNumber int8, lastNumber int8) int8 {
	result := firstNumber + lastNumber
	return result
}

func greetings2(firstName string, secondName string) (string, string) {
	greetingFirstName := "Hello " + firstName
	greetingSecondName := "Hello " + secondName
	return greetingFirstName, greetingSecondName
}

func greetings3() (firstName, middleName, lastName string) {
	firstName = "Akmal"
	middleName = "Keisin"
	lastName = "Alfateh"

	return firstName, middleName, lastName
}

func summary2(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func sayHelloWithFilter2(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		name = "..."
	}

	return name
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}
