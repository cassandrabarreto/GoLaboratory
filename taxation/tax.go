package main

import (
	"fmt"
	"taxation/filemanager"
	"taxation/prices"
)

func main(){
	taxRates := []float64{0,0.7,0.1,0.15}
	doneChannels := make([]chan bool, len(taxRates))
	errorChannels := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChannels[index] = make(chan bool)
		errorChannels[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPrice(fm, taxRate)
		go priceJob.Process(doneChannels[index], errorChannels[index])
	}

	for _, doneChan := range doneChannels{
		<- doneChan
	}

	for index := range taxRates{
		select{
		case err := <- errorChannels[index]:
			if err != nil {
				fmt.Println("ERROR")
			}
		case <- doneChannels[index]:
			fmt.Println("DONE")
		}
	}
} 