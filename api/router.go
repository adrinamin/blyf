package api

import (
    "net/http"
    "fmt"
    "github.com/adrinamin/blyf/api/handlers"
)

func RegisterRoutes() {
    fmt.Println("Register routes.")
    http.HandleFunc("/blyf", handlers.GetFilesHandler)
    http.HandleFunc("/upload", handlers.UploadFileHandler)
}
