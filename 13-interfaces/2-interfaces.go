package main

import (
	"fmt"
	"log"
)

type user struct {
	name string
}

// Write implements io.Writer.
func (u user) Write(p []byte) (n int, err error) {
	// panic("unimplemented")
	fmt.Println(u)
	return n,err
}

func main() {
	u := user{name: "ajay"}
	l := log.New(u, "sales:", log.Lshortfile) //go into New and you will see the doSomething function like func 
											  //that we did in 1-interfaces, the first out io.Writer has Writer 
											  //interface where the method signature for method to implement
											  // using struct is available, use that to create a method like in 
											  //line 13, this will work the program
	l.Println()
}
// Polymorphism means that a piece of code changes its behavior depending on the
//concrete data it’s operating on // Tom Kurtz, Basic inventor

// "Don’t design with interfaces, discover them". - Rob Pike
// Bigger the interface weaker the abstraction // Rob Pike

//SOLUTION AS PASTED IN THE CHAT

// package main

 

// import (

//     "fmt"

//     "log"

// )

 

// type user struct {

//     name string

// }

 

// func (u user) Write(p []byte) (n int, err error) {

//     fmt.Println("user struct :", u.name)

//     return n, err

// }

 

// func main() {

//     u := user{name: "ajay"}

//     l := log.New(u, "sales:", log.Lshortfile)

//     l.Println()

// }

 