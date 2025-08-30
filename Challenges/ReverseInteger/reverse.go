package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main(){
	reverse(-123)

}

func reverse(x int) int {
	
	var isNegativeNumber = false

	if x < 0{
		isNegativeNumber = true
		x = -x
	}
	converted := strconv.Itoa(x)

	list := strings.Split(converted, "")

	slices.Reverse(list)
	reversed_string := strings.Join(list, "")
	
	reversed_number , err := strconv.Atoi(reversed_string)
	
	if err != nil {
		panic(err)
	}

	if isNegativeNumber {
		reversed_number = -(reversed_number)
		fmt.Println(reversed_number)
	}
	
	// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
	//-2147483648 through 2147483647
	if reversed_number < -2147483648 || reversed_number > 2147483647 {
		return 0
	}else{
		return reversed_number
	}

}