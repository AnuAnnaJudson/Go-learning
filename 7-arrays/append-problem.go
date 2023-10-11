package main

import "fmt"

func main() {
	// a := []int{10, 20, 30, 40, 50, 60}

	// b := a[2:5]

	// b = append(b, 999)

	// fmt.Println(a)
	// fmt.Println(b)
	a:=[]int{10,20,30,40,50}
	i:=a[1:5]
	fmt.Println(cap(i))
	fmt.Printf("%p\n",a)//gives location (index)
	a=append(a, 60)
	fmt.Printf("%p\n",a)
	i=append(i, 60)
	i[0]=999
	fmt.Println(a,cap(a),len(a))
	fmt.Println(i,cap(i),len(i))//notice the capacity of i have changed , hence ew array, and does not change a

	b:=a[1:5]
	b[0]=800
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(i)
	// https://go.dev/ref/spec#Appending_and_copying_slices
	/*
		append func working

		i := []int{10, 20, 30, 40, 50 } // len = 5 , cap =5
		append(i,60) // not enough cap so allocation is going to happen

	//  sufficiently large underlying array.
		underlying array -> [10 20 30 40 50,60,{},{}] len =6 cap = 8

	append(i,70,90,300) // allocation would happen as we don't have enough cap to fit three values
		underlying array -> [10 20 30 40 50,60,70,80,90, , , , ] len =9 cap = 13

		If the capacity of s is not large enough to fit the additional values, append allocates a new, sufficiently large underlying array that fits both the existing slice elements and the additional values.
		Otherwise, append re-uses the underlying array.
	*/

}