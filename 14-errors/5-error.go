package main

import (
    "errors"
    "log"
)
type FileError struct{
	Path string
	Dirname string
	Err error
}

func (q *FileError) Error() string {
    return q.Path+" file in " + q.Dirname + " directory : "+ q.Err.Error()
}

var ErrFileNotFound = errors.New("not found")

func main() {
    err := openFile("abc.txt", "www")
    if err != nil {
       log.Println(err)
       return
    }
}

func openFile(fileName string, dirName string)  error {
	return &FileError{
		Path: fileName,
		Dirname: dirName,
		Err: ErrFileNotFound,
	}
}