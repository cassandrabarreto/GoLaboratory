/*

Cancellable worker
A worker generates numbers every 500ms and prints them. But main cancels it after 2 seconds.
Constraints:

Create a context with cancel in main using context.WithCancel
Launch a goroutine that loops forever, and on each iteration either:

Prints the next number if no cancellation has happened
Stops if ctx.Done() fires


Use select to listen to both cases
After 2 seconds, main calls cancel() to stop the worker
Use time.After(500 * time.Millisecond) as a ticker to control the speed of the loop

*/

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context){

	for {
		select {
		case <- ctx.Done():
			fmt.Println("CANCELLED")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Working")
		}
	}
}

func main(){
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(2 * time.Second)
	cancel()
}