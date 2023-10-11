package main

import (
	"errors"
	"fmt"
)
type FileError struct{
	Path string
	Dirname string
	Err error
}

func (q *FileError) Error() string { //method for error interface,note pointer of struct
    return q.Path+" file in " + q.Dirname + " directory : "+ q.Err.Error() //string how the errror should look like
}

var ErrFileNotFound = errors.New("not found") //custom error

func main() {
	var q *FileError
    err := openFile("abc.txt", "www")
    // if err != nil {
    //    log.Println(err)
    //    return
    // }
	if err != nil {
		if errors.As(err, &q) {
		   fmt.Println("custom error found in the chain", q.Path)
		   return
		}
		return
	}
}

func openFile(fileName string, dirName string)  error {
	return &FileError{ //return the pointer 
		Path: fileName,
		Dirname: dirName,
		Err: ErrFileNotFound,
	}
}