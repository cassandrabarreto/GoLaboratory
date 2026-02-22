package main

import "fmt"

// Write a function that finds and returns the maximum value in an array of 6 integers.
func findMax(arr [6]int) int {
    return 2
}

func main() {
    numbers := [6]int{45, 23, 67, 12, 89, 34}
	
	for i := 0; i < len(numbers)-1; i++ {
		if (numbers[i] > numbers[i +1]){
			println(numbers[i])
		}
		
	}

    result := findMax(numbers)
    fmt.Println(result)
}