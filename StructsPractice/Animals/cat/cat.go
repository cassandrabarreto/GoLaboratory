package cat

import (
	"errors"
	"fmt"
)

/*Create struct*/
type Cat struct {
	name string
	age string
}

/*Constructor function : Function that creates a struct instance*/
func NewCat(name, age string) (*Cat, error){
	
	if name == "" {
		return nil, errors.New("dog name must exist")
	}

	return &Cat{
		name : name,
		age: age,
		}, nil
}

func (catData *Cat) OutputCatData() {
	fmt.Println(catData.name, catData.age)
}

func (catData *Cat) MakeSound() error {
	fmt.Println("Meow!")
	return nil 
}

func (catData *Cat) Move() error {
	fmt.Println(catData.name + " just jumped!")
	return nil
}


