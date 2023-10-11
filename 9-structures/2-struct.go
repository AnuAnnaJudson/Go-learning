package main

import "fmt"

type User struct {
	name  string
	age   int
	marks []int
}

func main() {
	u := User{
		name: "John", 
		age: 25, 
		marks: []int{10, 20, 30},
	}
	fmt.Println(u.name, u.age,u.marks)
	fmt.Printf("%+v\n", u) // extra information of type along with what is assigned to what
	fmt.Printf("%v\n", u)  //shows just the content assigned to each of the struct variable attributes alone

}
