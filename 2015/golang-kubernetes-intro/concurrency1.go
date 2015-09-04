package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		sayHi()
	}
}

func sayHi() {
	time.Sleep(50 * time.Millisecond)
	fmt.Println("Hi")
}
