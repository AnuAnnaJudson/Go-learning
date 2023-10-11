[4:28 PM] Diwakar (Guest)

package main

import (
    "fmt"
    "sync"
)

func main() {
    ch := make(chan int)
    wg := sync.WaitGroup{}
    wgWorker := sync.WaitGroup{}
    wg.Add(1)
    //go1
    go func() {

       for i := 1; i <= 10; i++ {
          wgWorker.Add(1) // keeping track of number goroutines spawned
          go func(id int) {
             defer wgWorker.Done()
             ch <- id
          }(i)
       }
       wgWorker.Wait() // waiting until the worker goroutines are not finished
       close(ch)       //sending is finished over the channel ch
       wg.Done()
    }()

    //ranging until the channel is not closed
    for v := range ch {
       fmt.Println("recv", v)
    }
    wg.Wait()
}
 