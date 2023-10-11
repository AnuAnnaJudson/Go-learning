package mysql

import (
    "fmt"
    "project-3/stores"
)

type Conn struct {
    db string
}

func New(db string) Conn {
    return Conn{db: db}
}

func (m Conn) Create(usr stores.User) error {
    fmt.Println("adding to postgres", usr)
    return nil
}
func (m Conn) Update(usr stores.User) error {
    fmt.Println("updating in postgres", usr)
    return nil
}

func (m Conn) Delete(usr stores.User) error {
    fmt.Println("deleting in postgres", usr)
    return nil
}