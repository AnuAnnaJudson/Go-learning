package main

import (
	"net/http"
	"project-4/handlers"
)

func main(){
	http.HandleFunc("/home",handlers.Home)
    //start your server
    panic(http.ListenAndServe(":8080", nil))

}