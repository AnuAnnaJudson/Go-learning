package main

import "fmt"

type user struct {
    name string
}

func main() {
    var i interface{} = "str" //concept of empty interface says that you can add any type value to it and
							  //empty interfaces are also said to be implemented without methods automatically
    //var a any // any is an alias to interface{}
    i = 10
    i = true
    i = "hello"
    var u user //user is an empty struct
    i = u		//value of u assigned to i
    str, ok := i.(user) // type assertion, here the str will return value and ok is a bool ,
						// if it contains the type after . , of retruns true else false

    if !ok {
       fmt.Println("user type is not there")
       return
    }
    fmt.Println(str)
}