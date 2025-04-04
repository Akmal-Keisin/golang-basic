package main

func main() {
	var randomResult any = random()
	println(randomResult)

	checkType(randomResult)

	var randomResultString string = random().(string)
	println(randomResultString)

	var randomResultInt int = random().(int)
	println(randomResultInt) // panic: interface conversion: interface {} is string, not int

	var randomResultBool bool = random().(bool)
	println(randomResultBool) // panic: interface conversion: interface {} is string, not bool
}

func random() interface{} {
	return "Helllo World"
	// return 123
	// return true
}

func checkType(data interface{}) {
	switch data.(type) {
	case string:
		println("string")
	case int:
		println("int")
	case bool:
		println("bool")
	default:
		println("unknown")
	}
}