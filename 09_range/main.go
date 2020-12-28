package main

import "fmt"

func main() {
	//RANGE
	fmt.Println("RANGE")
	ids := []int{1, 2, 3, 4, 5, 6, 7, 832, 234, 345, 456}

	fmt.Println(ids)
	for i, id := range ids {
		fmt.Printf("%d - Id: %d\n", i, id)
	}
	//NOT USING INDEX
	for _, id := range ids {
		fmt.Printf("Id: %d\n", id)
	}

	// Adding all Ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("The sum is:", sum)

	//MAP LOOPING
	emails := map[string]string{"Tom": "tomh386@gmail.com", "Katie": "akh6122@yahoo.com", "Bob": "bob@bob.com"}
	for k, v := range emails {
		fmt.Printf("%s has an email of %s\n", k, v)
	}
	
}
