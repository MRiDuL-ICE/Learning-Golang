package main

import "fmt"

func main() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Orange")
	fmt.Println(fruitList)

	fruitList = append(fruitList[2:3])

	fmt.Println(fruitList)
}
