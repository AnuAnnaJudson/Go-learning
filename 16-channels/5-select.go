package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)
    c3 := make(chan string)
    done := make(chan struct{})//empty struct has no memory, also a channel created just for signalling
    wg := sync.WaitGroup{}//for the worker that is for seperate channels recieving the chanel
    wgWorker := sync.WaitGroup{}//this waitgroup is for signalling

    //wgWorker keep track of if the go routine work is finished
    //or not and we wil close the channel when work is done

    wgWorker.Add(3)//worker eait group is 3 as 3 channels are sending
    go func() {
       defer wgWorker.Done()
       time.Sleep(4 * time.Second)
       c1 <- "1"
       c1 <- "4"
    }()
    go func() {
       defer wgWorker.Done()
       time.Sleep(2 * time.Second)
       c2 <- "2"
    }()
    go func() {
       defer wgWorker.Done()
       c3 <- "3"
    }()

    wg.Add(1)//adds 1 to counter for done channel waitgroup
    go func() {//done chanel go routine
       defer wg.Done() //this is related to the done channel and will only work after the wgworker wait that is when recieving is complete for the worker chanels
       wgWorker.Wait() //waits for the worker chanel to finish
       close(done) // close the channel when goroutines are finished sending, also acts as a signal which signals the select to stop when the worker waitgroups are done working/sending
                  //this is also where the signalling happens or sending to done chanel happens
    }()

    wg.Add(1) //this is for the select  go routine
    go func() {
       defer wg.Done() //decrement after select is finished 
       for {
          // whichever case is not blocking exec that first
          //whichever case is ready first exec that.
          //case chan recv , send , default
          select {
          case x := <-c1://reciver for ch1
             fmt.Println(x)

          case x := <-c2://reciver for ch2
             fmt.Println(x)

          case x := <-c3://reciver for ch3
             fmt.Println(x)
          case <-done: // this case will exec when channel is closed, when done chanel is signalled this works
             fmt.Println("work is finished")
             return

          }
       }
    }()
    wg.Wait()//this is for the outer wait group to finish exec b4 exiting from the main go routine
    //fmt.Println(<-c1)
    //fmt.Println(<-c2)
    //fmt.Println(<-c3)

}
 