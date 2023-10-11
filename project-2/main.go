package main

import (
	"fmt"
	"go-training/project-2/db"
	"log"
)

func main(){
	result,ok:=db.NewConn("my database")
	if ok{
		fmt.Println("Connected Successfully",result)
	}else{
		fmt.Println("Connection unsuccessful",result)
	}
}