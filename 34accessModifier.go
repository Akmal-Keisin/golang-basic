package main

import (
	"basic/helper"
	"fmt"
)

func main() {
	appName := helper.AppName // Accessing exported variable
	fmt.Println("App Name:", appName)

	// appVersion := helper.Version // Accessing unexported variable (will cause an error)
	// fmt.Println("App Version:", appVersion)

	helper.AccessingSayGoodbye("John") // Accessing unexported function within the same package
	fmt.Println(helper.AccessingVersion()) // Accessing unexported variable within the same package
}