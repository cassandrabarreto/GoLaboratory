package main

import (
	"slices"
	"strconv"
	"strings"
)

func main(){
	reverse(-123)

}

func reverse(x int) int {
	
	var value int = x
	var isNegativeNumber = false

	converted := strconv.Itoa(value)

	list := strings.Split(converted, "")

	if list[0] == "-"{
		list = slices.Delete(list,0, 1)
		isNegativeNumber = true
	}

	slices.Reverse(list)
	reversed_string := strings.Join(list, "")
	
	reversed_number , err := strconv.Atoi(reversed_string)
	
	if err != nil {
		panic(err)
	}

	if isNegativeNumber {
		reversed_number = -(reversed_number)
	}
	
	// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
	//-2147483648 through 2147483647
	if reversed_number < -2147483648 || reversed_number > 2147483647 {
		return 0
	}else{
		return reversed_number
	}

}