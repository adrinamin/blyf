package main

import (
	"fmt"
	"github.com/adrinamin/blyf/api"
	"io/fs"
	"net/http"
	"os"
)

const (
	FilePath = "files"
)

func main() {

	fmt.Println("Initializing....")
	fmt.Println("Creating upload dir")
	err := os.Mkdir(FilePath, fs.ModeDir)
	if err != nil && !os.IsExist(err) {
		fmt.Println(err)
	}

	// register
	api.RegisterRoutes()

	// start http server
	fmt.Println("Starting server on port 8080.")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
