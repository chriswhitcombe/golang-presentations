package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go sayHi()
	}
	// Hack!
	time.Sleep(100 * time.Millisecond)
}

func sayHi() {
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Hi")
}
