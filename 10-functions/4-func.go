package main

import "fmt"

type funcType func(endpoint string)

func get(endpoint string) {
	fmt.Println("This is a Get Request")
}

func post(endpoint string) {
	fmt.Println("This is a Post Request")
}

func put(endpoint string) {
	fmt.Println("This is a Put Request")
}

func request(op funcType,endpointName string) {
	op("something")
	fmt.Println(endpointName)
}

func main() {
	request(put,"wow")
}