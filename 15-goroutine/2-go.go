package main

import (
    "fmt"

    "sync"
)

func main() {
    wg := &sync.WaitGroup{}
    // counter to add number of goroutines
    // that we are starting or spinning up
	//this in turn ensures that the main program doesn't exit before all the goroutines 
	//have completed their tasks.
    wg.Add(1)
    go hello(wg) // spawning a goroutine
	// wg.Add(1) // should be there for each go routine
	// go bye(wg)
    fmt.Println("end of the main")

    wg.Wait() //block until the counter is reset to 0
	//have to it only then give output , but have to have jsut one of this atleast in main
    // bad way of doing it
    //time.Sleep(2 * time.Second)
}

func hello(wg *sync.WaitGroup) {
    defer wg.Done() //decrement the counter
	//should be there for each go routine
	//defer executes when this function returns and executes no matter what
    fmt.Println("hello from the hello func")
}

// func bye(wg *sync.WaitGroup){
	// defer wg.Done()
    // fmt.Println("Bye from the Bye func")

// }