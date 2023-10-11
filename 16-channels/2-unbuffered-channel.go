package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	ch := make(chan int) //buffered channel

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 20
				//using waitgroups here gives error as there is 2 senders and only
				// one reciever for unbuffered channel
		ch <- 10

	}()

	x := <-ch //it is a blocking call until we don't recv the value
	// x=<-ch
	y := <-ch
	fmt.Println(x, y)

	fmt.Println("end of main")
	wg.Wait()

}
