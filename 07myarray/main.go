package main

import "fmt"

func main() {
	fmt.Println("Exploring array in go")

	// here array length initialize is mandatory
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Grape"
	fruitList[2] = "Orange"
	fruitList[3] = "Banana"

	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("Fruit list is : ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("Veg list is : ", vegList)
	fmt.Println("Veg list is : ", len(vegList))
}
