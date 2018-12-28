package main

import "fmt"

func main() {
	// var name = "Brad"
	// Short Hand
	name := "Brad"
	var age = 37
	const isCool = true
	size := 1.3

	// name, email := "wzh", "wzhkobe@gmail.com"

	fmt.Println(name, age, isCool)
	// Get Type
	fmt.Printf("%T\n", size)
}
