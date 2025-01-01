package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    fmt.Println("Hello, Podman!")

    // todo: implement upload and download functionality
    // todo: implement a web server which handles http requests for upload, download and delete

    message := "Welcome to blyf"
    
    // definition
    handler := func (w http.ResponseWriter, req *http.Request) {
        io.WriteString(w, fmt.Sprintf("%s\n", message))
    }
    
    // register
    http.HandleFunc("/blyf", handler)

    // start http server
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }

}
