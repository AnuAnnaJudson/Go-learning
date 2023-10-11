package main

import "fmt"

type InfoLogger interface {
	info() string
}

type zap struct {
	path string
}

type loggus struct {
	path string
}

func (z zap) info() string {
	fmt.Println("String from zap method" + z.path)
	return "String from zap method" + z.path
}
func (l loggus) info() string {
	fmt.Println("String from loggus method" + l.path)
	return "String from loggus method" + l.path
}

func doLog(I InfoLogger) {
	I.info()
}

func main() {
	m := zap{
		path: "---path of zap---",
	}
	n := loggus{
		path: "---path of loggus---",
	}
	doLog(m)
	doLog(n)

}

//SOLUTION FROM CHAT
// package main

 

// import "fmt"

 

// type Infoo interface {

//     Info()

// }

 

// type zap struct {

//     terminal string

// }

 

// func (z zap) Info() {

//     fmt.Println(z.terminal)

 

// }

 

// type logress struct {

//     file string

// }

 

// func (l logress) Info() {

//     fmt.Println(l.file)

// }

 

// func Dolog(i Infoo) {

//     i.Info()

// }

 

// func main() {

//     //z1 := zap{"zip.com"}

//     z2 := logress{"filee.com"}

//     Dolog(z2)

 

// }