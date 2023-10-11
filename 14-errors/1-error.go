package main

import (
    "errors"
    "fmt"
    "log"
)

var user = make(map[int]string)

// Error string should not be capitalized or end with punctuation mark

var ErrRecordNotFound = errors.New("record not found")
var ErrStingEmpty=errors.New("string is empty")
func FetchRecord(id int) (string, error) {
    
    name, ok := user[id]
    if !ok {
       return "", ErrRecordNotFound
    }
    return name, nil
}

func FetchData(s string)(string,error){
	if s==""{
		return s,ErrStingEmpty
	}
	return s,nil
}

func main() {

    // n, err := FetchRecord(10)
    // if err != nil {
    //    log.Println(err)
    //    return
    // }
    // fmt.Println(n)
	str,err:=FetchData("ooooooooooooo")
	if err!=nil{
		log.Println(err)
		return
	}
	fmt.Println(str)
}