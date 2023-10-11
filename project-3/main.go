package main

import (
	"project-3/stores"
	"project-3/stores/mysql"
	"project-3/stores/postgres"
)

func main() {
	u := stores.User{
		Name:  "ajay",
		Email: "ajay@email.com",
	}
	m := mysql.New("mysql") //as db string in conn cannot be accessed outside mysql, created a new func to 
							//initialize it mysql:packagename New:function name
	// stores.StorerInterface = m //assigning mysql instance to storer
							   // we can do this because mysql impls all the methods
							   // of the storer interface
	// stores.StorerInterface.Create(u)

	pg:=postgres.New("postgres")
	// stores.StorerInterface=pg
	// stores.StorerInterface.Update(u)
	ms := stores.NewService(m) //this is called from the stores.go file where we initialize the instance 
							   //of the interface using the instance of type mysql
	ms.Create(u)

	ps := stores.NewService(pg) //this is called from the stores.go file where we initialize the instance 
								//of the interface using the instance of type postgres
	ps.Create(u)
}