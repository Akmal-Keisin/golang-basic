package main

import "fmt"

// - int8	8 bits/1 byte	-128 to 127
// - int16	16 bits/2 byte	-32768 to 32767
// - int32	32 bits/4 byte	-2147483648 to 2147483647
// - int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807

func main() {
	fmt.Println("Hello Number")

	// This is using int or alias of int32
	fmt.Println("One: ", 1)
	fmt.Println("Two: ", 2)

	// This is using float
	fmt.Println("Three point four: ", 3.4)
}
