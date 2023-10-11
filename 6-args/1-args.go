package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// fmt.Println(os.Args)
	// fmt.Printf("%T", os.Args)
	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	a:=os.Args[1:]
	fmt.Println(a)
	if len(a)!=3{
		log.Println("Inside if block")//log gives date and time
		
		return //do not have to use else 
	}

	fmt.Println("Outside if block")

}
