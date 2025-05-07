package main

import (
	"fmt"
)

func main(){
	fmt.Println("Variadic function app. ")
	fmt.Println(sum(1,1,4))
}

/* A varidiac function ccepts a variable number of arguments of the same type.*/
func sum(nums ...int) int{
	total := 0
	for _, num := range nums {
	    total += num
	}
	return total
}