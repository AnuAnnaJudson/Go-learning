// [12:12 PM] QUAZI, MINHAJUDDIN (CW)
// main.go
package main
 
import (
    "errors"
    "fmt"
    "net/http"
)
 
func main() {
    url := "httmkp://google.com"
    a := &api{
        httpClient: &http.Client{},
    }
    sc, _ := a.MakeAPI(url)
    fmt.Println("sc:", sc)
}
 
type HTTPClient interface {
    Do(req *http.Request) (*http.Response, error)
}
 
type api struct {
    httpClient HTTPClient
}
 
func (a *api) MakeAPI(url string) (statusCode int, err error) {
    httpRequest, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        statusCode = http.StatusInternalServerError
        err = errors.New(http.StatusText(http.StatusInternalServerError))
 
        return
    }
 
    httpResponse, err := a.httpClient.Do(httpRequest)
    if err != nil {
        statusCode = http.StatusInternalServerError
        err = errors.New(http.StatusText(http.StatusInternalServerError))
 
        return
    }
 
    statusCode = httpResponse.StatusCode
 
    if statusCode == http.StatusUnauthorized {
        err = errors.New(http.StatusText(http.StatusUnauthorized))
 
        return
    }
 
    return
}
 
// [12:12 PM] QUAZI, MINHAJUDDIN (CW)
