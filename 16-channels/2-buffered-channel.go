package main

 

import (

    "fmt"

)

 

func main() {

 

    ch := make(chan int, 1) //buffered channel

 

    go func() {

       ch <- 20

       ch <- 10

    }()

 

    x := <-ch //it is a blocking call until we don't recv the value
	x=<-ch
    fmt.Println(x)

 

    fmt.Println("end of main")

 

}