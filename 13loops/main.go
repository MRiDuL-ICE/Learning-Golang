package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for idx, day := range days {
		fmt.Printf("index is %v and value is %v \n", idx, day)
	}

	rougueValue := 1

	for rougueValue <= 10 {

		if rougueValue == 2 {
			goto devOps
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}

devOps:
	fmt.Println("Diving into DevOps")
}
