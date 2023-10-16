package main

import (
	"fmt"
	"sync"
)

var myMap = make(map[int]int)
var wg sync.WaitGroup

func main() {
	var m = &sync.Mutex{}
	for i := 0; i <= 2; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			m.Lock()
			myMap[n] = n * n
			m.Unlock()

		}(i)
	}
	fmt.Println(myMap)

	wg.Wait()
}
