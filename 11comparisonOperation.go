package main

import "fmt"

func main() {
	// ==	Equal to	x == y
	// !=	Not equal	x != y
	// >	Greater than	x > y
	// <	Less than	x < y
	// >=	Greater than or equal to	x >= y
	// <=	Less than or equal to	x <= y

	var value1 string = "Akmal"
	var value2 string = "Keisin"

	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

	var value3 int8 = 25
	var value4 int8 = 20

	fmt.Println(value3 == value4)
	fmt.Println(value3 != value4)
	fmt.Println(value3 > value4)
	fmt.Println(value3 < value4)
	fmt.Println(value3 >= value4)
	fmt.Println(value3 <= value4)
}
