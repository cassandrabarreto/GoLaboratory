package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main(){
	addBinary("11", "1")
}

func addBinary(a string, b string) string {

	// Transform Numbers
	big_number_1 := new(big.Int)
	big_number_1.SetString(a,2)

	big_number_2 := new(big.Int)
	big_number_2.SetString(b,2)
	
	// Create big Sum 
	sum := new(big.Int)

	// Sum numbers
	result := sum.Add(big_number_1, big_number_2)
	fmt.Println(result)

	// Convert it to string
	binary_string := fmt.Sprintf("%b", result)
	println(binary_string)

	// Convert binary string into list
	binary_characters := strings.Split(binary_string, "")
	fmt.Println(binary_characters)

	// Convert to String
	converted_string := strings.Join(binary_characters, "")

	// Remove trailing zeros
	processed_binary_string := strings.TrimLeft(converted_string, "0")
	
	if processed_binary_string == "" {
		processed_binary_string = "0"
	}
	fmt.Println("Number"+ processed_binary_string)
	return processed_binary_string
	 

}