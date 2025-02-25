package main

import "fmt"

func main() {

	// Declare an array with 3 string item
	var shoppingList [3]string
	shoppingList[0] = "Apple"
	shoppingList[1] = "Avocado"
	shoppingList[2] = "Oil"

	// Can't define item with another type except from the declaration
	// shoppingList[2] = 10

	// Can't define item more than the declaration
	// shoppingList[3] = "Oil"

	fmt.Println(shoppingList)

	// Another way to declare an array
	var oddNumber = [3]int{
		1,
		3,
		5,
	}

	fmt.Println(oddNumber)

	// Default value of array
	var defaultValueInt = [10]int{}
	fmt.Println(defaultValueInt)

	var defaultValueString = [10]string{}
	fmt.Println(defaultValueString)

	// Function array

	// Get array length
	fmt.Println(len(shoppingList))

	// Get array item by index
	fmt.Println(shoppingList[1])

	// Update array item by index
	shoppingList[1] = "Banana"
	fmt.Println(shoppingList[1])

	// Inferred length array
	var dynamicLengthArray = [...]int{
		10,
		12,
		13,
		10,
	}
	fmt.Println(dynamicLengthArray)

	// Can't define a inferred length array without define the inside first,
	// the length of the array will be set to 0
	// var dynamicLengthArray2 = [...]int{}
	// dynamicLengthArray2[1] = 12
	// fmt.Println(dynamicLengthArray2)
	// error
}
