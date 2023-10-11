package main

import "fmt"

func main() {
	a := make([]int, 0, 5)// 1st no. is length and 2nd no. is capacity
	//"make" can also mean plan in advance in terms of cap and len
	/*
	[12:59 PM] QUAZI, MINHAJUDDIN (CW)

	you have started you length from 2, 
	so at the beginning it will add default values in the beginning which is 0 for integers. 
	the new elements will be added after the defined length, ie after first two 0 elements

	if you want to avoid the zeros you are getting in the beginning,
	give length as zero in you make func while creating a slice
	*/
	// fmt.Println(cap(a))
	// fmt.Println(len(a))
	fmt.Printf("%p\n",a)
	a=append(a, 100,200,300,400,500,600)
	fmt.Printf("%p\n",a)
	fmt.Println(a)
}