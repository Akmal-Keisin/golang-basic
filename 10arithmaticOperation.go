package main

import "fmt"

func main() {

	// +	Addition	Adds together two values	x + y
	// -	Subtraction	Subtracts one value from another	x - y
	// *	Multiplication	Multiplies two values	x * y
	// /	Division	Divides one value by another	x / y
	// %	Modulus	Returns the division remainder	x % y

	var number1 = 10
	var number2 = 2

	var result1 = number1 + number2
	var result2 = number1 - number2
	var result3 = number1 * number2
	var result4 = number1 / number2
	var result5 = number1 % number2

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Augmented assigment
	// =	x = 5	x = 5
	// +=	x += 3	x = x + 3
	// -=	x -= 3	x = x - 3
	// *=	x *= 3	x = x * 3
	// %=	x %= 3	x = x % 3

	result6 = 12
	result6 += number1
	fmt.Println(result6)

	result6 -= number1
	fmt.Println(result6)

	result6 *= number1 * number2
	fmt.Println(result6)

	result6 /= number1 / number2
	fmt.Println(result6)

	result6 %= number1 % number2
	fmt.Println(result6)

	// Unary Arithmatic Operator
	number1++
	fmt.Println(number1)
	number2--
	fmt.Println(number1)
}
