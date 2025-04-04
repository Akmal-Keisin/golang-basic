package main

import "fmt"

func main() {
	var data map[string]string = newMap("Keisin");
	if data != nil {
		fmt.Println(data);
		} else {
		fmt.Println("Empty");
	}
}

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}