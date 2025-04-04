package main

import "fmt"

type Address struct {
	city     string
	province string
	country  string
}

func main() {
	var address *Address = new(Address)
	var address2 *Address = address

	address.city = "Semarang"
	address.province = "Jawa Tengah"
	address.country = "Indonesia"

	fmt.Println(address)
	fmt.Println(address2)
}