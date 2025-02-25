package main

import "fmt"

func main() {

	// Crreate a map datatype
	var person map[string]string = map[string]string{}
	person["name"] = "Akmal Keisin Alfateh"
	person["address"] = "Gedawang"

	// Another way to create map
	person2 := map[string]string{
		"name":    "Akmal Keisin Alfateh",
		"address": "Gedawang",
	}

	// Output: map[address:Gedawang name:Akmal Keisin Alfateh]
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// Output:
	fmt.Println(person["phone_number"])

	// Output: map[address:Gedawang name:Akmal Keisin Alfateh]
	fmt.Println(person2)
	fmt.Println(person2["name"])
	fmt.Println(person2["address"])

	// Another way to make map
	var book map[string]string = make(map[string]string)
	book["title"] = "Bumi"
	book["author"] = "Tere Liye"
	book["penerbit"] = "Gramedia"
	book["wrong_key"] = "Wrong Value"

	// Function Map

	// Get map length
	fmt.Println(len(book))

	// Get value from a specific key
	fmt.Println(book["title"])

	// Update value from a specific key
	book["wrong_key"] = "Right Key"
	fmt.Println(book)

	// Delete key & value with a key
	delete(book, "wrong_key")
	fmt.Println(book)
}
