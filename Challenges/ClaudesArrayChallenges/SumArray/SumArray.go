package main

import "fmt"

// Write a function that takes an array of 5 integers and returns their sum.

func sumArray(arr [5]int) int {
    return sumNumbers(arr, 0)
    
}

func sumNumbers(numbersToSum [5]int, index int) int{
	if index >= len(numbersToSum){
		return 0
	}

	return numbersToSum[index] + sumNumbers(numbersToSum, index+1)
}

func main() {
    numbers := [5]int{10, 20, 30, 40, 50}
    result := sumArray(numbers)
    fmt.Println(result)
}