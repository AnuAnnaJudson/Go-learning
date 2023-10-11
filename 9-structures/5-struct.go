package main

import "fmt"

type user struct {
	name string
	age  int
}

type movies struct {
	movieName string
	*user     // Embedding a pointer to the user struct
}

func (m *movies) getUserName() string {
	return m.name
}

func main() {
	u := user{
		name: "xyz",
		age:  29,
	}

	m := movies{
		movieName: "abc",
		user:      &u,
	}

	fmt.Println(m.movieName)
	fmt.Println(m.getUserName())
	fmt.Println(m.age)
}
