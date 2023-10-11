package main

import "fmt"

var User=make(map[int][]string) //When declaring a global variable a short hand wont 
								//work but the var declaration
								//note: you cannot use key with incomparable datatype, value can be of any datatype

func search(id int){
	s,ok:=User[id]//returns value to s and ok is a boolean variable that set to true if userid or key exist and false
					// for the reverse
	if ok{
		fmt.Println("user exist with value as: ",s)
		delete(User,id)

		return
	}
	fmt.Println("User id does not exist")

}

func main() {

	User[1]=[]string{"eat","drink","sleep"} //note you cannot simply asign values 
											//into slice usoing = use this syntax
	User[2]=[]string{"dance","sing","sleep"}
	User[3]=[]string{"sing","laugh","sleep"}
	search(1)


	for k, v := range  User{
		fmt.Println(k,":",v)
	}

}