package main

import "fmt"

func main() {
	dictionary := make(map[string]string,10)//map[key] value
	//unordered collection, order changes with each execution
	dictionary["up"] = "above"
	dictionary["below"] = "down"

	fmt.Println(dictionary["up"])
	for k, v := range dictionary { //range visits each element from 0till the end
		fmt.Println(k,":", v)
	}
	fmt.Println(len(dictionary))

}
