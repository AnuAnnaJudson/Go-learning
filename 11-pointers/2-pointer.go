package main

import (
	"log"
	"os"
)

func main() {

	// l := log.Lshortfile
	// log.Println(l)
	l:=log.New(os.Stdout,"sdn",log.Lshortfile)
	l.Println("something")
}