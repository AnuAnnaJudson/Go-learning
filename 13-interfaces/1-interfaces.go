package main

import "fmt"

// Polymorphism means that a piece of code changes its behavior depending on the
//concrete data it’s operating on // Tom Kurtz, Basic inventor

type Speaker interface {
	Speak()
}
type human struct {
	name string
}

func (h human) Speak() {
	fmt.Println("human speaking", h.name)
}

type ai struct {
	name string
}

func (a ai) Speak() {
	fmt.Println("ai speaking", a.name)
}

func doSomething(s Speaker) {
	fmt.Printf("%T\n", s)
	s.Speak()
}

func main() {
	h := human{name: "dev"}
	a := ai{name: "alexa"}
	doSomething(h)
	doSomething(a)
}
