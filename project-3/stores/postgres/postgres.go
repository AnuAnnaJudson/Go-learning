package postgres

import "fmt"
import "project-3/stores"

type Conn struct {
	db string
}

//CHAT PASTED SOLUTION
func New(db string) Conn {
    return Conn{db: db}
}
func (p Conn) Create(usr stores.User) error {
    fmt.Println("adding to postgres", usr)
    return nil
}
func (p Conn) Update(usr stores.User) error {
    fmt.Println("updating in postgres", usr)
    return nil
}

func (p Conn) Delete(usr stores.User) error {
    fmt.Println("deleting in postgres", usr)
    return nil
}


// func New(db string) Conn {
//     return Conn{db: db}
// }

// func (c Conn) Create(usr stores.User) error {
// 	fmt.Println("Create User")
// 	return nil

// }
// func (c Conn) Update(usr stores.User) error {
// 	fmt.Println("Update User")
// 	return nil

// }
// func (c Conn) Delete(usr stores.User) error {
// 	fmt.Println("Delete User")
// 	return nil
// }
 
