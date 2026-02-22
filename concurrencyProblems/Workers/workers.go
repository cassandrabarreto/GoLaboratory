package main

import (
	"fmt"
	"sync"
)

func work(id int, wg *sync.WaitGroup, jobsChan <- chan int){
	defer wg.Done()
	for job := range jobsChan {
		fmt.Printf("worker %d handled job %d", id, job)
	}
}
	func main(){
		jobsChan := make(chan int)
		
		var wg sync.WaitGroup

		wg.Add(3)
		for i:= 1; i <=3; i++{
			go work(i, &wg, jobsChan)
		}

		for i:= 1; i <=5; i++{
			jobsChan <- i
			
		}
		close(jobsChan)
		wg.Wait()
		fmt.Println("Done")
}