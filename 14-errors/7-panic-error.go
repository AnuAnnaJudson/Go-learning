package main

import (
	"fmt"
	// "runtime/debug"
)

func main() {
	//defer recoveryFunc()//it will only run after the main statements are over, that means if damage 
						  //done by others, after that
    doSomething()
    fmt.Println("end of the main, stopped gracefully")
}

func doSomething() {
    defer recoveryFunc()//always put recover function in defer which means it run no matter what
    var i []int
    i[100] = 1000
}

func recoveryFunc() {
    msg := recover()// recover is an inbuilt function that recovers the program even when panic happens
					//recover func // it recovers from the panic and stop panic further propagation
    if msg != nil { // nil means no panic // otherwise msg var would be the msg of the panic
		// fmt.Println(string(debug.Stack()))
       fmt.Println(msg)
       fmt.Println("recovered from the panic")
    }
}