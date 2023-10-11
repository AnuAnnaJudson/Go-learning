package main

import "fmt"

func main() {
	i := 10
	p := &i //adrress of i is given to p
	//p := &i // var p *int
	//fmt.Println(p)
	//fmt.Println(&i)
	update(p) //passing the value of p
	fmt.Println(i)
	fmt.Println(&p, "p") //content of p
}
func update(ptr *int) {
	*ptr = 10000
	fmt.Println(&ptr, "ptr") //content of ptr
}
