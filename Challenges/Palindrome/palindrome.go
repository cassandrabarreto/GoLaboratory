package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main(){
	isPalindrome(-121)
}

func isPalindrome(x int) bool {
    
	// Int to String
	original_number_text := strconv.Itoa(x)
	
	// Reverse String logic
	reversed_list := strings.Split(original_number_text, "")
	slices.Reverse(reversed_list)


	processed_reversed_text := strings.Join(reversed_list, "")

	fmt.Println("Original Number:", original_number_text)
	fmt.Println("Reversed Number:", processed_reversed_text)
	
	if original_number_text == processed_reversed_text {
		return true
	}else{
		return false
	}

}