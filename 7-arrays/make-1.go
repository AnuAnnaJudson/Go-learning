package main

import (
	"fmt"
	"os"
)

func main() {
	a := []int{10, 20, 30, 40, 50}
	b := make([]int, len(a[1:3]), 100) //depending on the length here the number of elements copied from a changes
	//if 0, 0 elements will be copied from a to b and son on
	copy(b, a[1:3]) //1st argument is destination n 2nd source
	//copy creates deep copy that is sources's backing array to destination's backing array

	// b[0]=100
	b = append(b, 100)
	fmt.Println(a)
	fmt.Println(b)
    fmt.Printf("%T",os.Args)// returns []string which means it a slice of the type string

	X:=[]string{"A","B","C","D","E"}
	Y:=make([]string,len(X))
	// Y=X //here it will change X also as Y is just a reference
	copy(Y,X) //here it will not change X  as Y is copy of X
	Y[0]="AAA"
	fmt.Println(X)
	fmt.Println(Y)


}
