package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
func sayHi() {
	fmt.Println("Hi")
}

*/

/*
func sayHiSlowly(phrase string, doneChan chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println(phrase)
	doneChan <- true
}*/

func sendNumber(OutChannel chan int){
	for{
		number := rand.Int()
		time.Sleep(500  * time.Millisecond)
		OutChannel <- number
	}
}
	
func main(){
	numChan := make(chan int)
	go sendNumber(numChan)

	timeOutChan := time.After(2 * time.Second)
	
	for {
		select{
		case num := <- numChan:
			fmt.Println(num)
		case <- timeOutChan:
			fmt.Println("TIMEOUT")
			return
	}
	}
} 