package main

import (
	"fmt"
	"math/rand/v2"
)

func main(){
	address := obtainVariableMemoryAddress(2)
	fmt.Printf("The memory address of the provided variable is: %p \n", address)

	modifiedNumber :=  modifyVariableValueToRandomValue(344)
	fmt.Printf("Your transformed int is: %d\n", modifiedNumber)

	fmt.Printf(checkIfPointersPointToSameAddress())

}

func obtainVariableMemoryAddress(number int)*int{
	myNumber := number
	var p *int 
	p = &myNumber
	return p
}

func modifyVariableValueToRandomValue(number int) *int {
	myNumber := number
	p := &myNumber
	*p = int(rand.Int32())
	return p
}


func checkIfPointersPointToSameAddress()string{
	number:= 2 
	pointer1 := &number
	pointer2 := &number

	if pointer1 == pointer2 {
		return "Pointers are pointing to the same address."
	} else {
		return "Pointers are not pointing towards the same address."
	}
}