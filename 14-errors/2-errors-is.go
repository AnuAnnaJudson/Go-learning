package main

 

import (

    "errors"

    "fmt"

    "log"

    "os"

)

 

var ErrFileNotFound = errors.New("not in the root directory")
/*Error strings should not be capitalized (unless beginning with an exported name, 
a proper noun or an acronym) and should not end with punctuation. This is because error 
strings usually appear within other context before being printed to the user.*/
 

func main() {

    _, err := openFile("test.txt")

    if err != nil {

 

        log.Println(err)

        if errors.Is(err, ErrFileNotFound) { //checks if custom error in chain or not

            fmt.Println("File not found")

        }

        return

    }

}

 

func openFile(fileName string) (*os.File, error) {

    f, err := os.Open(fileName)

    if err != nil {

        return nil, fmt.Errorf("%v %w", err, ErrFileNotFound)

    }
    return f, nil

}