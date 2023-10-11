package main

import "fmt"

type user struct {
	name string
}


func main() {
	checkType(user{
		name:"king",
	})
}

func checkType(i interface{}) {//i any gives error coz go version is old so use interface{} which means
							   // any is an interface{}
	switch i.(type) {
	case int:
		fmt.Println("it is an int value")
	case user:
		fmt.Println("it is an user struct value")
	default:
		fmt.Println("it is fefault case")

	}
}
