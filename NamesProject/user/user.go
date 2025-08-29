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

/* This struct will 'embedd' User struct */
type Admin struct {
	email string
	password string
	/* user User can also be implemented but it's annoying to always call admin.User to use user functions */
	User
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

/* Create constructor for Admin struct*/
func NewAdmin(email string, password string ) *Admin {

	return &Admin{
		email: email,
		password : password,
		User: User {
			firstName : "ADMIN",
			lastName : "ADMIN",
			birthDate : "---",
			createdAt : time.Now(),
		}, 
	}
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
