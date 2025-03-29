package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	mridul := User{"MRiDuL", "abdulwahab22400@gmail.com", true, 25}

	fmt.Println(mridul)

	fmt.Printf("Details of mridul: %+v \n", mridul)

	mridul.GetStatus()
	mridul.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "tdsxmridul@gmail.com"

	fmt.Println("Email of this user: ", u.Email)
}
