package main

import (
	"context"
	"fmt"

	"sync"

	"time"
)

func main() {

	wg := new(sync.WaitGroup)

	// gives an empty container to put context values
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch := make(chan int)

	wg.Add(1)
	go func() {

		defer wg.Done()

		select {
		case ch <- 20:
			fmt.Println("values sent")
		default:
			fmt.Println("receiver not present")
		}

	}()
	recv(ctx, ch)
	wg.Wait()

}

func recv(ctx context.Context, ch chan int) {
	time.Sleep(time.Second)
	select {

	case val := <-ch:
		fmt.Println(val)

	case rc := <-ctx.Done():
		fmt.Println("cancelled", rc)

	}

	//this line

}
