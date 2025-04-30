package main

import (
	"Animals/cat"
	"Animals/dog"
	"fmt"
)

/*Sound Interface*/
type Sound interface {
	MakeSound() error
}

/*Embedded Interface*/
type Animal interface{
	Sound
	Move()
}

func main (){
	fmt.Println("Animal App")

	dogName := getAnimalData("Please enter your dog's name:")
	dogBreed := getAnimalData("Please enter your dog's breed:")
	dogAge := getAnimalData("Please enter your dog's age:")

	catName := getAnimalData("Please enter your cat's name:")
	catAge := getAnimalData("Please enter your cat's age:")

	appDog, err := dog.NewDog(dogName, dogBreed, dogAge)
	if err != nil {
		fmt.Println("Dog error:", err)
		return
	}

	appCat, err := cat.NewCat(catName, catAge)
	if err != nil {
		fmt.Println("Cat error:", err)
		return
	}

	appDog.OutputDogData()
	appCat.OutputCatData()

	appDog.MakeSound()
	appCat.MakeSound()

}
	


func getAnimalData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

