package main

import "fmt"

type languages struct {
	lang string
}

func (u *languages) updateLanguage(newLang string) {
	u.lang = newLang
}

func main() {
	lang1 := languages{
		lang: "C++",
	}
	fmt.Printf("%+v",lang1)
	lang1.updateLanguage("GoLang")
	fmt.Printf("%+v ",lang1)
}