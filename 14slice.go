package main

import "fmt"

func main() {
	// Create slice from an Array
	// var myarray = [length]datatype{values} An array
	// myslice := myarray[start:end] A slice made from the array

	cars := [...]string{"Ferrari", "Lamborgini", "Supra", "Honda Civic", "Pagani", "Kijang", "toyota"}
	var slice1 []string = cars[2:4]
	fmt.Println(slice1)

	slice2 := cars[:3]
	fmt.Println(slice2)

	var slice3 []string = cars[3:]
	fmt.Println(slice3)

	var slice4 []string = cars[:]
	fmt.Println(slice4)

	// Function Slice
	// len(slice) to get the slice length
	fmt.Println(len(slice3))

	// cap(slice) to get the sllice capacity
	fmt.Println(cap(slice3))

	// Append Slice
	days := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	daySlice1 := days[:4]
	daySlice1[0] = "New Sunday"
	daySlice1[1] = "New Monday"
	fmt.Println(days)      // [New Sunday New Monday Tuesday Wednesday Thursday Friday Saturday]
	fmt.Println(daySlice1) // [New Sunday New Monday Tuesday Wednesday]

	daySlice2 := append(daySlice1, "New Holiday")
	fmt.Println(daySlice2) // [New Sunday New Monday Tuesday Wednesday New Holiday]

	daySlice2[0] = "Test"
	fmt.Println(daySlice2) // [Test New Monday Tuesday Wednesday New Holiday]
	fmt.Println(days)      // [Test New Monday Tuesday Wednesday New Holiday Friday Saturday]

	daySlice2 = append(daySlice2, "New Holiday 2")
	fmt.Println(daySlice2)

	// Create a slice
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Akmal"
	newSlice[1] = "Keisin"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	var newSlice2 []string = append(newSlice, "Alfateh")
	newSlice2[0] = "Akmal Edited"
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// Copy a slice
	var fromSlice []string = days[:]
	var toSlice []string = make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// Array vs Slice
	var thisIsArray = [...]int8{1, 2, 3, 4, 5, 6}
	var thisIsSlice = []int8{1, 2, 3, 4, 5, 6}

	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)
}
