package main

import "fmt"

func main() {
	// var fruitArr [2]string

	// fruitArr[0] = "apple"
	// fruitArr[1] = "orange"

	// fmt.Println(fruitArr)

	// declare and assign

	// fruitArr := [2]string{"Apple", "Orange"}
	// fmt.Println(fruitArr)

	fruitSlice := []string{"Apple", "Orange", "Banana"}

	fmt.Println(fruitSlice)

	fmt.Println(len(fruitSlice))

	fmt.Println(fruitSlice[1:2])
}
