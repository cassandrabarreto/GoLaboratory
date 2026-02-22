package main

import "fmt"

func main(){
	fmt.Println(sumNumbers(2,4))
}

func sumNumbers (num1 int, num2 int)(int){
	result  := num1 + num2
	return result
}
