package main

import "fmt"

func main()  {
	//POINTERS
	fmt.Println("Pointers")
	 a:= 5
	 b := &a

	 fmt.Println(a, b)
	 fmt.Println(a, *b)

	 c := a + *b
	 fmt.Println(c)

	 //Change val with pointer
	 *b = 23
	 fmt.Println(a)
}