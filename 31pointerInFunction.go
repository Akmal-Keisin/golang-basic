package main

import "fmt"

type Address struct {
	city     string
	province string
	country  string
}

func main() {
	var address Address = Address{
		city:     "Semarang",
		province: "Jawa Tengah",
		country:  "Indonesia",
	}

	fmt.Println("Before update")
	fmt.Println(address)
	updateCountryToIndonesia(&address)

	fmt.Println("After update")
	fmt.Println(address)
}

func updateCountryToIndonesia(address *Address) {
	address.country = "Indonesia"
}