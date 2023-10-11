package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50, 60}
	b := a[1:4]
	fmt.Println(b)
	b[0] = 777
	fmt.Println(b)
	fmt.Println(a)

}