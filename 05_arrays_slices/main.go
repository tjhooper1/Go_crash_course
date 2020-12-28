package main

import "fmt"

func main()  {
	// Arrays

	var fruitArray [2]string
	// Assign Values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	fmt.Println(fruitArray)

	// DECLARE AND ASSIGN
	newFruitArr := [2]string{"pineapples", "grapes"}
	fmt.Println(newFruitArr)

	// SCLIES
	fruitSlice := []string{"oranges", "strawberries", "blueberries"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:2]) // strawberries
}