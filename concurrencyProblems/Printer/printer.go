/*
Delayed printer
Launch 3 goroutines using anonymous functions. Each one should wait a different amount of time before printing:

First one waits 1 second then prints "one"
Second one waits 2 seconds then prints "two"
Third one waits 3 seconds then prints "three"

Constraints:

Use go func() { ... }() for each goroutine
Use wg.Go() to manage them
Main should wait for all 3 to finish before printing "done"
*/

package main

import (
	"fmt"
	"sync"
	"time"
)


func main(){
	var wg  sync.WaitGroup

	wg.Go(func(){
		time.Sleep(1 * time.Second)
		fmt.Println("one")
	})

	wg.Go(func(){
		time.Sleep(2 * time.Second)
		fmt.Println("two")
	})

	wg.Go(func(){
		time.Sleep(3 * time.Second)
		fmt.Println("three")
	})


	wg.Wait()
	fmt.Println("Done")

}