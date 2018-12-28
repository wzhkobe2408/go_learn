package main

import "fmt"

func main() {
	// define a map
	// emails := make(map[string]string)

	// // asign kv
	// emails["Bod"] = "bob@gmail.com"
	// emails["sharon"] = "sharon@gmail.com"

	// fmt.Println(emails)

	// delete(emails, "Bod")

	// fmt.Println(emails)

	// declare and asign
	emails := map[string]string{
		"Bod":  "bob@gmail.com",
		"mike": "mike@gmail.com",
		"kery": "kery@gmail.com",
	}

	fmt.Println(emails)
}
