package main

type validationError struct {
	Message string
}

func (validationError *validationError) Error() string {
	return validationError.Message
}

type notFoundError struct {
	Message string
}

func (notFoundError *notFoundError) Error() string {
	return notFoundError.Message
}

func saveData(id string, data any) error {
	// Simulate a validation error
	if id == "" {
		return &validationError{Message: "ID cannot be empty."}
	}

	// Simulate a not found error
	if id != "123" {
		return &notFoundError{Message: "Data not found."}
	}

	// Simulate successful data saving
	return nil
}

func main() {
	err := saveData("", "some data")
	if err != nil {

		if validationErr, ok := err.(*validationError); ok {
			println("Validation Error:", validationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			println("Not Found Error:", notFoundErr.Error())
		} else {
			println("Unknown Error:", err.Error())
		}
	} else {
		println("Data saved successfully.")
	}
}