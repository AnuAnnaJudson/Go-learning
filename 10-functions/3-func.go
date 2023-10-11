package main

import "fmt"

type operation func(x, y int)int //here it is a type of function that accepts 2 arguments and returns int val

func add( a, b int) (int) {
	return a + b
}
func sub( a, b int) (int) {
	return a - b
}

func operate(op operation){//here the type of this function is a function and pass a function 
									//from main to this function
									//note that the signature of the functio passed to this should
									//be equal as mentioned here in this definition
	fmt.Println(op(30,10))

}

func main() {
	operate(add)//you can pass the function with the same signature as the one given in operate argument

	mul:=func(x,y int)int{ //note you only write function inside main like this as a variable
		return x*y
	}
	fmt.Println(mul(2,2))
	
}