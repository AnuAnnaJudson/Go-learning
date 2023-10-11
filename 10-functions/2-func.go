package main

import (
	"fmt"
	"strconv"
)

func Addition(a , b string) (int, error) {
	x,err:=strconv.Atoi(a)//convert string to integer
	if err!=nil{
		return 0,err
	}
	y,err:=strconv.Atoi(b)
	if err!=nil{
		return 0,err
	}

	return x+y,nil
}

func main() {
	result,err:=Addition("10", "io")
	fmt.Println(result,err)
}