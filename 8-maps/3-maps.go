package main

import "fmt"

func main() {

	m := map[[1]int]interface{}{}

	m[[1]int{1}] = [1]int{1}

	fmt.Println("m:", m)

}