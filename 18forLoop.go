package main

import "fmt"

func main() {

	// This is first way to do for loop
	counter := 1
	for counter <= 10 {
		fmt.Println("Loop : ", counter)
		counter++
	}

	// This is another way to do it
	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Loop2 : ", counter2)
	}

	// Loop through data
	phones := [5]string{"Infinix", "Nubia", "Xiaomi", "Samsung", "Iphone"}
	for counter3 := 0; counter3 < len(phones); counter3++ {
		fmt.Println("This is : ", phones[counter3])
	}

	// Loop through data using range
	for index, phone := range phones {
		fmt.Println("Index : ", index, " Phone : ", phone)
	}

	// Loop through data using range without index
	for _, phone := range phones {
		fmt.Println("Phone : ", phone)
	}

	// Break & Continue
	for i := 0; i <= 10; i++ {
		if i == 5 {
			fmt.Println("This is the break point")
			break
		}

		fmt.Println("Index : ", i)
	}

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Index : ", i)
	}
}
