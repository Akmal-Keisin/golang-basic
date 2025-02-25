package main

import "fmt"

func main() {
	type NIM string

	var myNIM NIM = "1111111111 "
	fmt.Println(myNIM)
	fmt.Println(NIM("2222222222"))
}
