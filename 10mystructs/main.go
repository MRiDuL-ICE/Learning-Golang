package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	mridul := User{"MRiDuL", "abdulwahab22400@gmail.com", true, 25}

	fmt.Println(mridul)

	fmt.Printf("Details of mridul: %+v \n", mridul)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
