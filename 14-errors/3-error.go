package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println(strconv.Atoi("abc"))
    fmt.Println(strconv.Atoi("xyz"))
    fmt.Println(strconv.ParseInt("efg", 10, 64))
    fmt.Println(strconv.ParseFloat("qwer", 64))
}