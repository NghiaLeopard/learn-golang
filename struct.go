package main

import (
	"fmt"

	"example.com/struct/user"
)

// struct is type value , not reference
// then when use it , func will copy value
// if adjust value in struct then must use pointer

type str string

func (text str) log() {
	fmt.Println(text)
}


func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthDate (MM/DD/YYYY): ")

	var appUser user.User
	var err error

	var admin user.Admin

	admin = user.NewAdmin("nghiabeo1605@gmail.com","xxxxxx")
	admin.PrintData()
	admin.ClearData()
	admin.PrintData()

	appUser,err = user.New(firstName,lastName,birthDate)

	if err != nil {
		fmt.Print(err)
	}

	appUser.PrintData()
	appUser.ClearData()
	appUser.PrintData()

    // printData(&appUser)

	var name str = "Beo"

	name.log()
}

// func printData(u *user) {
// 	fmt.Println(u.firstName,u.lastName,u.birthDate)
// }

// input: pointer to integer
func getUserData(promptText string) string {
	var value string
	fmt.Print(promptText)
	fmt.Scan(&value)

	return value
}