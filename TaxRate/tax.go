package main

import "fmt"

func main() {

	prices := []float64{10,20,30}
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	
	// This is a map creation. 
	result := make(map[float64][]float64) 
	// iterate over all taxRates
	for _, taxRate := range taxRates {
		// This is a slice creation that takes number of prices as lenght. 
		taxIncludedPrices := make([]float64, len(prices))
		// for each existing price,  add the price rate calculation to the slice.
		for priceIndex , price := range prices {
			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
		}
		/*calculate result by adding to a map (like a python dict)
		 the taxRate as a 'key' and the slice as a value*/
		result[taxRate] = taxIncludedPrices
	}

	fmt.Println((result))

}