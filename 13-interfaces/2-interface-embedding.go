package main

import "fmt"

type Walker interface {
	walk()
}

type Runner interface {
	run()
}
type WalkRunner interface {
	Walker
	Runner
}

type man struct{}

func (m man) walk() {
	fmt.Println("Walking......")
}
func (m man) run() {
	fmt.Println("Running......")
}

func main() {
	var m man
	var w Walker = m
	var r Runner = m
	var rw WalkRunner = m
	w.walk()
	r.run()
	rw.run()
	rw.walk()

}
