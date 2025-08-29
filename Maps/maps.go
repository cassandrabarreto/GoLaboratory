package main

import "fmt"


func main(){

	employees := map[string]int{
		"John" : 27,
		"Jane" : 27,
		"Anna" : 23,	
		"Troy" : 10000,
	}
	
	employees["Charles"] = 1 
	fmt.Println(employees)
	fmt.Println(employees["John"])

	delete(employees, "Anna")
	fmt.Println(employees)

	/* When wanting to use indexes you can use make*/
	// Helpful when you have a raw estimate of the number of items allocated in slice
	userNames := make([]string, 2, 5)
	userNames[0] = "Jules"

	
	fmt.Println(userNames)

	for index, value := range userNames{
		fmt.Println(index)
		fmt.Println(value)
	}
	

	
}