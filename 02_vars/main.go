package main

import "fmt"

var age int = 32 // var needed for global or out of main function

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128
	
	const isCool = true
	// isCool = false - cannot redefine
	name := "Tom" // shorthand
	var size float32 = 1.3

	favColor, favNumber := "blue", 9 // quick assign

	fmt.Println(name, age, isCool, favColor, favNumber)
	fmt.Printf("%T\n", size) // data type
}
