package main

import (
	"fmt"
	"project-1/models" //modeulename/package name nite that module name should be same as in go mod file

)

func main(){
	u:=models.Users{
		Name:"John",
		Email:"john@gmail.com",
		Password:"12345",
		Permissions:map[string]bool{ //note how a map is initailised
			"read":true,
			"write":false,
		},
		Admin:true,
	}
	fmt.Printf("%+v\n",u)
}