package main

import "fmt"

func main()  {
	x := 12
	y := 10

	if x < y {
		fmt.Println(x) // 5
	} else {
		fmt.Printf("%d is less than %d\n", y, x) // 10 is less than 12
	}

	//else if
	color := "blue"
	if color == "red" {
		fmt.Println(color)
	} else if color == "blue" {
		fmt.Println(color)
	} else {
		fmt.Println("color not blue or red")
	}

	// SWITCH
	switch color {
	case "red":
		fmt.Println("the color is red")
	case "blue":
		fmt.Println("the color is blue")
	default:
		fmt.Println("the color is not red or blue")
	}
}