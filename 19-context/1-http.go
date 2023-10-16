package main

import (
    "context"
    "fmt"
    "io"
    "log"
    "net/http"
    "time"
)

func main() {

    // gives an empty container to put context values
    ctx := context.Background()
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()//either to cancel your work or to clear resources
    //constructing the request // we are not making the call
    r, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com/", nil)
    if err != nil {
       log.Println(err)
       return
    }

    //doing the request // DefaultClient have no timout
    b, err := http.DefaultClient.Do(r)
    if err != nil {
       log.Println(err)
       return

    }
    body, err := io.ReadAll(b.Body)
    if err != nil {
       log.Println(err)
       return

    }
    fmt.Println(string(body))
}