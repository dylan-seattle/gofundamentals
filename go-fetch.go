package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    
    )
func getPage(url string) (int, error) {
    resp , err := http.Get(url)
    if err != nil {
        return 0, err 
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return 0, err 
    }
    return len(body), nil 
}
func getter(url string, size chan int) {
    length, err := getPage(url)
    if err == nil {
    size <- length
    }
}
func main() {
    urls := []string{"http://www.google.com", "http://www.yahoo.com", "http://reddit.com"}
    
    size := make(chan int)
    
    for _, url := range urls {
        go getter(url,size)
    } 
    for i := 0; i < len(urls); i ++ {
        fmt.Printf("Length %d\n", <- size)
    }
    
}