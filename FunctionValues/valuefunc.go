package main

import "fmt"


type transformFn func(int) int

func main(){
	numbers := []int{1,2,3,4}
	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	// Anonymous function example
	f := func(x int) int {
		return x * 2
	}
	fmt.Println(f(5))


	//fmt.Println(doubled)
}

// Function that takes another function as argument
func transformNumbers(numbers *[]int, transform transformFn) []int{
	dNumbers := []int{}

	for _, val := range *numbers{
		dNumbers = append(dNumbers,transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

