package main

import (
	"fmt"
	"sync"
)

func produce(jobs chan <- int){
	for i:= 0; i <= 6; i++{
		jobs <- i
	}
	close(jobs)
}

func worker(jobs <- chan int, results chan <- int){
	for i := range jobs{
		result := i * i
		results  <- result
	}
}

func main(){
	var wg sync.WaitGroup
	
	jobsChan := make(chan int)
	resultChan := make(chan int)

	go produce(jobsChan)

	wg.Go(func() {
		worker(jobsChan, resultChan)
	})
	
	go func() {
        wg.Wait()
        close(resultChan)
    }()

	for i := range resultChan{
		fmt.Println(i)
	}
	
}
