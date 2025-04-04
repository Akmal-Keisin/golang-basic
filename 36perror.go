package main

import (
	"errors"
)

func main() {
	result, err := division(10, 0)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Result:", result)
	}

	result, err = division(10, 2)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Result:", result)
	}
}

func division(firstValue int, secondValue int) (int, error) {
	if secondValue == 0 {
		return 0, errors.New("division by zero")
	}
	return firstValue / secondValue, nil
}