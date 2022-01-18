package main

import (
	"fmt"
	"sync"
	"time"
)

// sync asks for the order of desired execution in order to manage function execution

var wg sync.WaitGroup

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("PANIC")
	}
}

func printStuff() {
	// after done running the code and threads, tell wait group things are done
	defer wg.Done()

	// this is for error handling in case wg.Wait() takes too long because of an error
	defer handlePanic()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}
}

// func say(s string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(s)
// 		time.Sleep(time.Millisecond * 300)
// 	}
// }

// goroutine allows the code to be non-blocking
// but if every single function is set to a goroutine, nothing is going to execute before the goroutine finishes execution

func main() {
	// go say("Hello")
	// say("There")

	// wait group takes the Add function, takes an integer, add it to a behind-the-scenes queue
	wg.Add(1)
	go printStuff()

	// waits for execution of code and threads, wait for all counters to pop off the stack
	wg.Wait()

	// in order for goroutines to communicate with one another, we open up a channel of communication between goroutines to pass info back and forth, and do more complex logic
}
