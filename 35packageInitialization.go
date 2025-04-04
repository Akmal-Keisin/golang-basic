package main

import (
	"basic/database"
	_ "basic/internal"
	"fmt"
)

func main() {
	fmt.Println("Database Connection:", database.GetConnection()) // Accessing exported function
}