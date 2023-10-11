package main

import (
    "errors"
    "fmt"
    "log"
)

var ErrAdmission = errors.New("admission error")
var ErrFees = errors.New("fees error")
var ErrFinance = errors.New("finance not accepted")

// end user
func main() {
    err := admission()
    if err != nil {
       log.Println(err)
    }

}

func admission() error {
    err := fees()
    if err != nil {
       return fmt.Errorf("%w %w", err, ErrAdmission)// multiple %w will only work with latest go version
    }
    return nil

}

func fees() error {
    err := finance()
    if err != nil {
       return fmt.Errorf("%w %w", err, ErrFees)
    }
    return nil
}

func finance() error {
    
    return fmt.Errorf("%w", ErrFinance)

}