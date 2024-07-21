package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func NewAdmin(email,password string) Admin {
	
	return Admin{
		email: email,
		password: password,
		User: User{
		firstName: "Admin",
		lastName:  "Sub admin",
		birthDate: "16/05/2003",
		createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthDate string) (User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return User{}, errors.New("first name or last name or birth day is empty")
	}

	return User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

// add method in struct
func (u User) PrintData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearData() {
	u.firstName = ""
	u.lastName = ""
}