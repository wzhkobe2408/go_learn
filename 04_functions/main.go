package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2, int) int {
	return num1 + num2
}

// func getSum(num1, num2 int) int {

// }

func main() {
	fmt.Println("Hello World")
	fmt.Println(greeting("Brad"))
}
