package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    // register
    http.HandleFunc("/blyf", basicMessage)
    http.HandleFunc("/upload", uploadHandler)
    http.HandleFunc("/download/", downloadHandler)
    http.HandleFunc("/delete/", deleteHandler)

    // start http server
    fmt.Println("Starting server on port 8080.")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}

func basicMessage (w http.ResponseWriter, req *http.Request) {
    message := "Welcome to blyf"
   
    io.WriteString(w, fmt.Sprintf("%s\n", message))
}

func uploadHandler (w http.ResponseWriter, req *http.Request) {
    fileUploadMessage := "File upload"
    fmt.Println("Uploading file...")
    io.WriteString(w, fmt.Sprintf("%s\n", fileUploadMessage))
}

func downloadHandler (w http.ResponseWriter, req *http.Request) {
    fileDownloadMessage := "File download"
    fmt.Println("Download file...")
    io.WriteString(w, fmt.Sprintf("%s\n", fileDownloadMessage))
}

func deleteHandler (w http.ResponseWriter, req *http.Request) {
    fileDeleteMessage := "File delete"
    fmt.Println("Delete file...")
    io.WriteString(w, fmt.Sprintf("%s\n", fileDeleteMessage))
}
