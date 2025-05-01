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

	/*Output*/
	appDog.OutputDogData()
	appCat.OutputCatData()

	/*Make Sounds*/
	appDog.MakeSound()
	appCat.MakeSound()

	/*Move*/
	appDog.Move()
	appCat.Move()

}
	
func getAnimalData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

/*Interface that allows any value*/
func printAnything(value interface{}){

	/*Instead of adding the switch you can do the following:*/
	intVal , ok := value.(int)
	if ok {
		fmt.Println("It's an integer!", intVal)
		return 
	}

	floatVal , ok := value.(float64)
	if ok {
		fmt.Println("It's a float!", floatVal)
		return 
	}

	stringVal , ok := value.(string)
	if ok {
		fmt.Println("It's a string!", stringVal)
		return 
	}

	/*Add special switch that checks type*/
	// switch value.(type) {
	// case int:
	//	fmt.Println("It's an integer!")
	//case string:
	//	fmt.Println("It's a string!")
	//case float64:
	//	fmt.Println("It's a float!")
	//}	
}


/*Generic function that takes any int or float value*/
func add[T int|float64](a, b T) T{
	return a + b
}

