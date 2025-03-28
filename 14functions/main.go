package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Functions in golang")

	proRes, myMsg := proAdder(4, 7, 6, 4)

	fmt.Println(proRes, myMsg)
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi pro result for proAdder"
}

func greeter() {
	fmt.Println("Hello, from golang")
}
