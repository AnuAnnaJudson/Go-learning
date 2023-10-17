package main

import "fmt"

func add(s func(a, b int) int) int {
	fmt.Println("Add")
	s(1, 2)
	return 1
}
func sub(a, b int) int {
	fmt.Println("Evaluate Sub", b-a)
	return 1
}

func operate(op func(s func(a, b int) int) int) {
	fmt.Println("operate")
	op(sub)
	// return
}

func main() {
	operate(add)
}
