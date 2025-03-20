package main

import "fmt"

func main() {

	// // basic example of pointer
	// var ptr *int
	// fmt.Println("value of my pointer is", ptr)

	myNumber := 44

	// refrence a pointer to a value
	var ptr = &myNumber

	// will print the memory referenece
	fmt.Println("Value of actual pointer is ", ptr)

	// will print the value of pointer
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr * 2

	fmt.Println("My new value is ", myNumber)
}
