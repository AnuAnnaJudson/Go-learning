package main
import "fmt"

type user struct{
	name string
	age int
}

var record []user

func (u user) addUsersToRecord(){
	record = append(record, u)
}

func main(){
	u1:=user{
		name:"John",
		age:9,
	}
	u2:=user{
		name:"jooo",
		age:90,
	}
	u3:=user{
		name:"yon",
		age:67,
	}
	u1.addUsersToRecord()
	u2.addUsersToRecord()
	u3.addUsersToRecord()

	// users:=[]user{
	// 	u1,u2,u3,
	// }
	// fmt.Println(users)
	for i,v:=range record{//here i is index and v is value
		fmt.Println(i,v)
	}
	// fmt.Printf("%T",record)
}