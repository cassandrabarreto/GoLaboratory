package main

import (
	"fmt"
	"time"
)

func sendFast(FastChannel chan string){
	message := "fast"
	time.Sleep(1 * time.Second)
	FastChannel <- message
}

func sendSlow(SlowChannel chan string){
	message := "slow"
	time.Sleep(3 * time.Second)
	SlowChannel <- message
}

func main(){
	FastChan := make(chan string)
	go sendFast(FastChan)
	
	SlowChan := make(chan string)
	go sendSlow(SlowChan)

	select{
	case <- SlowChan:
		fmt.Println("WIN")
		return
			
	case <- FastChan:
		fmt.Println("WIN")
		return
	}
}