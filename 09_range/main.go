package main

import "fmt"

func main() {
	ids := []int{33, 345, 768, 87, 32, 546}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not use index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("sum", sum)

	// range with map
	emails := map[string]string{
		"Bod":  "bob@gmail.com",
		"mike": "mike@gmail.com",
		"kery": "kery@gmail.com",
	}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

}
