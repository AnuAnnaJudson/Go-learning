package main

import (
	"fmt"

)

type money int //underying type here is int and money is the new type

func main() {
	var rupee money = 100
	fmt.Printf("%T",rupee)
	var i int64
	fmt.Println(i)

}