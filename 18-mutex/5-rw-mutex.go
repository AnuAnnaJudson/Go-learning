package main

import "sync"

type cabs struct {
	drivers int
	rw      sync.RWMutex
}

func main() {

}
