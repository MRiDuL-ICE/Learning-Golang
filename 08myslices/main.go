package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Orange")
	fmt.Println(fruitList)

	fruitList = append(fruitList[2:3])

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 45
	highScores[1] = 98
	highScores[2] = 75
	highScores[3] = 65

	highScores = append(highScores, 35, 57, 49, 58)

	sort.Ints(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	fmt.Println(highScores)

}
