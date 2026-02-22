package main

import "fmt"

func main (){
	age := 32

	var agePointer *int

	agePointer = &age
	fmt.Println(*agePointer)


} 