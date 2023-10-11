package main

import "fmt"

type user struct {
	name     string
	password string
}

var records=make(map[int]user)

func (u user) addToRecords(userid int) {
	// fmt.Println(u)
	records[userid] = u
}

func main() {
	u1 := user{
		name:     "John",
		password: "12345",
	}
	u2 := user{
		name:     "Tom",
		password: "12345",
	}
	u1.addToRecords(1)
	u2.addToRecords(2)
	fmt.Printf("%+v\n",records)

}