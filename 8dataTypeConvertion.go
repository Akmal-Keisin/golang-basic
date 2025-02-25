package main

import "fmt"

// int	Depends on platform:
// 32 bits in 32 bit systems and
// 64 bit in 64 bit systems	-2147483648 to 2147483647 in 32 bit systems and
// -9223372036854775808 to 9223372036854775807 in 64 bit systems
// - int8	8 bits/1 byte	-128 to 127
// - int16	16 bits/2 byte	-32768 to 32767
// - int32	32 bits/4 byte	-2147483648 to 2147483647
// - int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807

func main() {

	// Number convertion
	var value32 int32 = 32768
	var value64 int64 = int64(value32)
	var value16 int16 = int16(value32)

	// Output: 32768
	fmt.Println(value64)

	// Output: 32768
	fmt.Println(value32)

	// Output: -32768
	// If the convertion is downgrade, then the number before convertion is larger than the maximum of the number type,
	// it will converted to the lowest number of the converted type. In this case, the int16 has maximum number of 32767
	// but the convertion number is 32768, the result of the convereted number will be -32768, if the convertion number is far
	// more than that like 40000, the result of the converted number will be -32768 (the lowest number) + (40000 - 32767 (the max number))
	// and the output will be
	fmt.Println(value16)

	// String convertion
	var name = "Akmal Keisin Alfateh"
	var firstLetter uint8 = name[0]             // converted to uint8 65
	var firstLetterString = string(firstLetter) // A

	fmt.Println(name)
	fmt.Println(firstLetter)
	fmt.Println(firstLetterString)

}
