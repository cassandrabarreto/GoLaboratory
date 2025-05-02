package main

import (
	"fmt"
)

type Product struct {
	title string
	id string
	price float64
}

func main(){
	// first way of creating an array
	prices := [4]float64{10.99, 9.99, 34.3, 11.1}
	
	// Dogs Slice. This i the recommended way of creating arrays.
	dogs := []string{"Bobby", "Marcus", "Lola"}

	//Second way of creating an array
	var productNames [4]string = [4]string{"Novel", "Magazine", "Dictionary"}

	productNames[3] = "Children Book"
	// The first value will be included. Second value will be excluded. 
	featuredPrices := productNames[0:2]
	fmt.Println(featuredPrices)

	// Include all the values expect but the one with index  3
	featuredPrices = productNames[:3]
	fmt.Println(featuredPrices)

	featuredPrices = productNames[1:4]

	fmt.Println(productNames[3])
	fmt.Println(featuredPrices)

	highlightedPrices := featuredPrices[:2]

	// Prints size of slice or array
	fmt.Println(len(highlightedPrices))

	// Prints capacity (elements in can hold) of slice or array
	fmt.Println(cap(highlightedPrices))

	printIndexValues(prices)

	// Append method to add elements to array. But you can only add slices to arrays.
	updatedPrices := append(dogs, "Slinky")
	fmt.Println(updatedPrices)
}

func printIndexValues(arr [4]float64){
	for index, element := range arr{
		fmt.Printf("%v has the index of number: %v\n", element, index)
	}
}
