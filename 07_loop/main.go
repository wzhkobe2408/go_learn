package main

import "fmt"

func main() {
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for loop
	// asign; condition; increatment

	// for ; i <= 20; i++ {
	// 	fmt.Println(i)
	// }

	// fizzbuzz
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
