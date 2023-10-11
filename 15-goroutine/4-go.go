package main

import (
    "fmt"
    "sync"
)

func main() {
    wg := &sync.WaitGroup{}
    //wg.Add(10)
    for i := 1; i <= 10; i++ {
       wg.Add(1)
       go func(a int) {//here i is a receptor or a parameter
		 //if you just use closure value of i without passing it then will give 11 as,
		 //the main will only spawn go routine will not run it until main excecution statements 
		 //are over,if you pass the i from outside, then the copy of the outer i is used for
		 //internal function thus giving all the expected values of i
          defer wg.Done()
          fmt.Println("work", a)
       }(i)//arguments are passed in a function call,thus making it from closure to pass by value 
	   	   //than reference
    }
	fmt.Println("End of main")
    wg.Wait()//when encounters forces to execute the go routines
}