package main

import (
	"errors"
	"fmt"
)

func main(){
	userName := getUserData("Please Type user's name:")
	userAge := getUserData("Please Type user's age:")
	userProfession := getUserData("Please Type user's profession:")

	cassandra , err := newUser(userName, userAge, userProfession)
	cassandra.outputUserData()

	if err != nil{
		fmt.Println(err)
		return
	}
}

type user struct{
	name string
	age string
	profession string
}

func newUser(name string, age string , profession string)(*user, error){
	if name == ""{
		return nil, errors.New("User must have a name.")
	}
	return &user{
		name : name,
		age : age,
		profession: profession,
	}, nil
}

func getUserData(prompt string) string{
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	return value
}

func (userData user) outputUserData() {
	fmt.Println(userData.age, userData.name,userData.profession)
}
