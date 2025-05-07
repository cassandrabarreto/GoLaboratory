package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to closures program.")

	// A Go closure is a function that captures variables from its surrounding scope
	// They are implemented through anon functions
	
	add := adder()

	fmt.Println(add(10))
	fmt.Println(add(10))
	fmt.Println(add(10))
}

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}
	
