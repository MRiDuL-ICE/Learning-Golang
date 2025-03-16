package main

import (
	"fmt"
)

func main() {
	var username string = "Mridul"
	var password int = 123456
	var isLoggedIn bool = true
	fmt.Println(username, password, isLoggedIn)
	fmt.Printf("Variable type of : %T \n", username)
	fmt.Printf("Variable type of : %T \n", password)
	fmt.Printf("Variable type of : %T ", isLoggedIn)
}
