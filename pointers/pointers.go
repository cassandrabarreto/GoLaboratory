package main

import "fmt"


func main(){
	age := 26
	agePointer := &age
	//adultYears := getAdultYears(age)
	//fmt.Println(adultYears)
	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int){
	//return *age - 21
	*age = *age - 18
}