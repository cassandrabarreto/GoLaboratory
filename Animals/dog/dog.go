package dog

import (
	"errors"
	"fmt"
)

/*Create struct*/
type Dog struct {
	name string
	breed string
	age string
}

/*Constructor function : Function that creates a struct instance*/
func NewDog(name, breed, age string) (*Dog, error){
	
	if name == "" {
		return nil, errors.New("dog name must exist")
	}

	return &Dog{
		name : name,
		breed : breed,
		age: age,
		}, nil
}

func (dogData *Dog) OutputDogData() {
	fmt.Println(dogData.name, dogData.breed, dogData.age)
}

func (dogData *Dog) MakeSound() error {
	fmt.Println("WOOOOOF!")
	return nil 
}
