package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func main() {
	var customer Customer
	customer.Name = "Akmal Keisin Alfateh"
	customer.Address = "Gedawang"
	customer.Age = 20

	var customer2 Customer = Customer{
		Name:    "Dr. Krinzz",
		Address: "Gedawang",
		Age:     20,
	}

	var customer3 Customer = Customer{"Person", "Semarang", 20}
	customer3.sayHello()

	fmt.Println(customer)
	fmt.Println(customer2)
	fmt.Println(customer3)

}
