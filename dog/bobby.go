package main

import (
	"errors"
	"fmt"
)

/*Create struct*/
type dog struct {
	name string
	breed string
	age string
}

/*Constructor function : Function that creates a struct instance*/
func newDog(name, breed, age string) (*dog, error){
	
	if name == "" {
		return nil, errors.New("dog name must exist.")
	}

	return &dog{
		name : name,
		breed : breed,
		age: age,
		}, nil
}

func main (){
	fmt.Println("Dog App")
	dogName := getDogData("Please enter your dog's name:")
	dogBreed := getDogData("Please enter your dog's breed:")
	dogAge := getDogData("Please enter your dog's bread:")

	appDog , err := newDog(dogName, dogBreed, dogAge)
	appDog.outputDogData()

	if err != nil{
		fmt.Println(err)
		return
	}
}

func getDogData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
func (dogData dog) outputDogData() {
	fmt.Println(dogData.name, dogData.breed, dogData.age)
}