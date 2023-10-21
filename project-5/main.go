package main

//run main.go then in cmd do curl localhost:8080/getuser or open this url in browser
import (
	"net/http"
	"project-5/handlers"
)

func main() {
	http.HandleFunc("/getuser", handlers.GetUser) //handlefunc is a type which accepts a handler function
	//register a pattern & name of func that handles the request are the 2 arguments accepted by a handlefunc
	//a handler func is the one which itself handles a request
	//start your server
	panic(http.ListenAndServe(":8080", nil))
}
