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

	var address2 *Address = &address

	address2.city = "Jakarta"
	address2.province = "DKI Jakarta"
	address2.country = "Indonesia"

	fmt.Println("Address 2")
	fmt.Println(address2)

	fmt.Println("Address 1")
	fmt.Println(address)

	// address2 = &Address{
	// 	city: 	"Bandung",
	// 	province: "Jawa Barat",
	// 	country:  "Indonesia",
	// }

	// fmt.Println("Address 2")
	// fmt.Println(address2)

	// fmt.Println("Address 1")
	// fmt.Println(address)

	*address2 = Address{
		city:     "Bandung",
		province: "Jawa Barat",
		country:  "Indonesia",
	}

	fmt.Println(address)
	fmt.Println(address2)
}