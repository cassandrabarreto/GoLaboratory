package main

import (
	"fmt"
	"sync"
)

func generate(numsChan chan <- int){
	for i := 1 ; i <= 5; i++{
		numsChan <- i
	}
	close(numsChan)
}

func double(numsChan <- chan int, doubleNumsChans chan <- int){
	for i := range(numsChan){
		number := i * 2
		doubleNumsChans <- number
	}
	close(doubleNumsChans)
}

func main (){
	var wg sync.WaitGroup

	numsChan := make(chan int)
	doubleChan := make(chan int)

	wg.Go(func() {
		generate(numsChan)
	})
	wg.Go(func() {
		double(numsChan, doubleChan)
	})

	for i := range doubleChan{
		fmt.Println(i)
	}
	wg.Wait()
	
}