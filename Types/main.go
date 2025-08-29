package main

import (
	"fmt"
)


type str string


func(text str) log() {
	fmt.Println(text)
}

/* You can call function using a custom type*/
func main() {
	var message str = "Hello"
	message.log()
}



