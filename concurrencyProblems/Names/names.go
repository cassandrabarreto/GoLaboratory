package main

import (
	"fmt"
	"sync"
)

func Greet(name string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Hi", name)
}

func main(){
	var wg sync.WaitGroup
	wg.Add(3)
	go Greet("Alice", &wg)
	go Greet("Bob", &wg)
	go Greet("Charlie", &wg)

	wg.Wait()
	fmt.Println("Done")

}