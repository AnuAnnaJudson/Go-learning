package main

import "fmt"

type user struct {
    name string
    age  int
}

type movies struct {
    movieName string
    // anonymous field which have no name
    // anonymous field get there names from the types
    user // embedding
}

type perms struct {
	user
    permission []string
}

type books struct {
	user
    bookName string
}

func (u *user) updateAge(newAge int){
	u.age=newAge
}

func (m *movies) getUserName() string{
	return m.name
}

func (m *movies) updateMovieName(newMovieName string){
	m.movieName=newMovieName
}



func main() {
    m := movies{
       movieName: "abc",
       user: user{
          name: "xyz",
          age:  29,
       },
    }

	book:=books{
		user:user{
			name:"king",
			age:30,
		},
		bookName:"Harry Potter",
	}
	m.updateAge(100)

    fmt.Println(book.bookName,m.age)
	fmt.Println(m.getUserName())
	fmt.Printf("%+v\n",m)
	m.updateMovieName("wooooooooooooooooooow")
	fmt.Printf("%+v\n",m)

}