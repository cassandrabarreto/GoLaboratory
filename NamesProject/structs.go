package main

import (
	"NamesProject/user"
	"fmt"
)


func main() {
	retrievedFirstName := getUserData("Please enter your first name: ")
	retrievedLastName := getUserData("Please enter your last name: ")
	retrievedBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	/* Create external struct instance */
	var appUser *user.User

	appUser , err := user.New(retrievedFirstName, retrievedLastName, retrievedBirthdate)

	if err != nil{
		fmt.Println(err)
		return
	}

	appUser.OutputUserData()
	appUser.ClearUserName()
	appUser.OutputUserData()

	/*  
	You can also use:
	appUser = user{
		firstName: retrievedFirstName,
		lastName: retrievedLastName,
		birthDate: retrievedBirthdate,
		createdAt: time.Now(),
	}
		*/
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}


