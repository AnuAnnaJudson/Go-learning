package main

import (
	"log"
	"sync"
	"time"
)

var counter = 0
 
func main() {
    wg := new(sync.WaitGroup)
	m:=new(sync.Mutex)
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
			m.Lock()//avoids raise condition, it waits for one routine to complete first
            j := counter
            time.Sleep(time.Millisecond)//line number 19 to 22 is the critical section as it tried to access a shared resource which is the counter
            j = j + 1
            counter = j
			m.Unlock()
        }()
    }
    wg.Wait()
    log.Println("counter:", counter)
}