package main

import (
	"fmt"
)


func main(){
	fact := factorial(3)
	fmt.Println(fact)
}

// Recursive function
func factorial(number int) int{
	if number == 0{
		return 1
	}
	return number * factorial(number - 1)
}