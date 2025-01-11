package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
	http.HandleFunc("/blyf", basicMessage)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download/", downloadHandler)
	http.HandleFunc("/delete/", deleteHandler)

	// start http server
	fmt.Println("Starting server on port 8080.")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func basicMessage(w http.ResponseWriter, req *http.Request) {
	message := "Welcome to blyf"

	io.WriteString(w, fmt.Sprintf("%s\n", message))
}

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// fileUploadMessage := "File upload"
	// io.WriteString(w, fmt.Sprintf("%s\n", fileUploadMessage))

	// Parse the multipart form, with a maximum memory of 10MB
	err := req.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		log.Println(err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	file, handler, err := req.FormFile("file")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// validate filename to avoid directory traversal attacks
	safeFileName := filepath.Base(handler.Filename)
	completeFileName := fmt.Sprintf("%s/%s", FilePath, safeFileName)
	exists := fileAlreadyExists(completeFileName)
	if exists {
        fmt.Fprintf(w, "File %s already exists. It will be renamed to %s + (1)\n", completeFileName, completeFileName)
		completeFileName = completeFileName + "(1)"
	}

	destinationFile, err := os.Create(completeFileName)
	if err != nil {
		log.Println("Error creating destination", err)
		http.Error(w, "Error creating destination file", http.StatusInternalServerError)
		return
	}

	defer destinationFile.Close()

	if _, err = io.Copy(destinationFile, file); err != nil {
		log.Println("Error saving destinationFile", err)
		http.Error(w, "Error saving destinationFile", http.StatusInternalServerError)
		return
	}

	// todo: Create a destination file
	// todo: Copy the uploaded file's content to the destination file

	fmt.Fprintf(w, "Upload of %s was successful.\n", handler.Filename)

}

func fileAlreadyExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func downloadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	path := req.URL.Path
	fmt.Printf("url path: %s\n", path)
	pathElements := strings.Split(path, "/")
	fileName := pathElements[len(pathElements)-1]
	fmt.Printf("File name: %s", fileName)
	io.WriteString(w, fmt.Sprintf("url path: %s\n", path))
	io.WriteString(w, fmt.Sprintf("file name: %s\n", fileName))
}

func deleteHandler(w http.ResponseWriter, req *http.Request) {
	fileDeleteMessage := "File delete"
	fmt.Println("Delete file...")
	io.WriteString(w, fmt.Sprintf("%s\n", fileDeleteMessage))
}
