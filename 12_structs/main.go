package main

import "fmt"

// Person - DEFINE PERSON STRUCT
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Methods for struct (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName
}

// Method for struct (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(s string) {
	if p.gender == "F" {
		p.lastName = s
	}
}

func main() {
	//structs
	fmt.Println("structs")
	// Init struct
	person1 := Person{firstName: "Tom", lastName: "Hooper", gender: "M", city: "Ormond Beach", age: 32}
	fmt.Println(person1)
	fmt.Println(person1.firstName)

	//alternative
	person2 := Person{"Katie", "Hooper", "Daytona Beach", "F", 34}
	fmt.Println(person2)
	fmt.Println(person2.greet())

	fmt.Println(person2.age)
	person2.hasBirthday()
	fmt.Println(person2.age)

	fmt.Println(person2.lastName)
	person2.getMarried("Williams")
	fmt.Println(person2.lastName)
}
