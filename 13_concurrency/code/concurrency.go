package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 300)
	}
}

// goroutine allows the code to be non-blocking
// but if every single function is set to a goroutine, nothing is going to execute before the goroutine finishes execution

func main() {
	go say("Hello")
	say("There")
}
