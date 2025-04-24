package user

import (
	"errors"
	"fmt"
	"time"
)

/* user starts in lowercase because it should not be accesible from outside. */
/* In this case it will be accessible by external packages */
type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

/* user starts in uppercase because it should be accesible from outside. */
func New(firstName, lastName, birthDate string) (*User, error) {

	if firstName == "" {
		return nil, errors.New("dog name must exist")
	}

	return &User{
	firstName : firstName,
	lastName : lastName,
	birthDate: birthDate,
	createdAt : time.Now(),
	}, nil
}

// This is how you specify a struct function 
func (userData *User) OutputUserData(){
	fmt.Println(userData.firstName, userData.lastName, userData.birthDate)
}

// Mutation Function
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
