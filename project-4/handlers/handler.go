package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Hobbies []string
	UserId  string
}

func Home(w http.ResponseWriter, r *http.Request) {
	//order is : 1. set content type 2. write status code 3. write response
	u := User{
		Hobbies: []string{"eat", "sleep", "run"},
		UserId:  "1",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	b, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(b)
	// w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(b))
}
