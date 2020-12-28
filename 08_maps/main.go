package main

import "fmt"

func main() {
	//MAPS
	fmt.Println("MAPS")

	// emails := make(map[string]string)

	// emails["Tom"] = "tomh386@gmail.com"
	// emails["Katie"] = "akh6122@yahoo.com"

	emails := map[string]string{"Tom": "tomh386@gmail.com", "Katie": "akh6122@yahoo.com"}

	println(emails)
	fmt.Println(emails["Tom"])
	fmt.Println(len(emails))

	//DELETE
	delete(emails, "Katie")
	println(emails)
}
