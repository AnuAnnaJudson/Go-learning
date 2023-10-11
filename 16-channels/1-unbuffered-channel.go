package main

import (
    "fmt"
)

func main() {

    ch := make(chan int) //unbuffered channel
                        //needs a reciever, if not will be i deadlock coz it keeps waiting
                        //https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a receiver is ready.
// send will block until there is no recv
 
                    

    go func() {

        ch <-10
       ch <- 20 // send will block until no receiver is ready
    }()

    x := <-ch //it is a blocking call until we don't recv the value
	//reciever gets ready first
    fmt.Println(x)

    fmt.Println("end of main")

}













// package main

// import "sync"

// func main() {
// 	wg := sync.WaitGroup{}
// 	wg.Add(1)
// 	a:=go work(wg)
// 	wg.Wait()
// }

// func work(wg *sync.WaitGroup) int {
// 	defer wg.Done()
// 	return 10
// }