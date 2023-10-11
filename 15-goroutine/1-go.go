package main

import (
	"fmt"
	"time"
)

func main() {
    go hello() // spawning a goroutine
	//main only gets the go routine ready
    fmt.Println("end of the main")
	time.Sleep(2*time.Second)
}

func hello() {
    fmt.Println("hello from the hello func")
	
}