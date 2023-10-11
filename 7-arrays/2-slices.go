package main

import "fmt"

func main() {
	var i []int //when no size is given atl its slice
	//i[1] = 100 // update
	i=append(i,20,30 ) //append grows your slice
	fmt.Println(i)
	if i == nil {
		fmt.Println("slice is nil", i)
	}
	
	b := []int{10, 230, 4050, 60}
	fmt.Println(b)
}