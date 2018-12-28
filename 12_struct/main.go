package main

import (
	"fmt"
	"strconv"
)

// Person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Person method
// greeting method (value receiver) dont change anything
func (p Person) greet() string {
	return "hello, my name is " + p.firstName + strconv.Itoa(p.age)
}

// pointer reciever
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	person := Person{
		firstName: "Kobe",
		lastName:  "Bryant",
		city:      "Boston",
		gender:    "male",
		age:       24,
	}
	fmt.Println(person)
	fmt.Println(person.firstName)
	person.age++
	fmt.Println(person)
	fmt.Println(person.greet())
}
