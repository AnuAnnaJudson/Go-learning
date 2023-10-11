package main

import (
    "fmt"
    "sync"
)

func main() {
    wg := &sync.WaitGroup{}
    //wg.Add(10)
    for i := 1; i <= 10; i++ {
       wg.Add(1)
       go work(i, wg)//10 go routines
    }
    wg.Wait() //its okay to just have wait here as the wg is passed by reference to work and hence wait wg
}

func work(i int, wg *sync.WaitGroup) {
    defer wg.Done() //should add when using wg.Add or will be in deadlock
    wg.Add(1)//since this is another go routine add counter value
    go func(id int) { //next 10 go routines
       defer wg.Done()//new set of go routine needs to decrement the counter
       fmt.Println("work", id)
    }(i)

}