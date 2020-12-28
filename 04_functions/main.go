package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(a int, b int) int {
	return a + b
}

func main()  {
	greet := greeting("Tom")
	sum := getSum(3, 5)

	fmt.Println(greet)
	fmt.Println(sum)
}