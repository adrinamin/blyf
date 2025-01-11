package main

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
    "github.com/adrinamin/blyf/api"
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
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Welcome to blyf!\n Here are your current files:\n ")

	files, err := ioutil.ReadDir(FilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for _, file := range files {
		fmt.Fprintf(w, "%s\n", file.Name())
	}

	w.WriteHeader(http.StatusOK)
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
		fmt.Fprintf(w, "File %s already exists. Please rename your file and try it again.\n", completeFileName)
		// completeFileName = completeFileName + "(1)"
		http.Error(w, "File upload failed because file name already exists.", http.StatusBadRequest)
		return
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
	w.WriteHeader(http.StatusCreated)
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
	fmt.Fprintf(w, "Trying to download content of file: %s/%s\n", FilePath, fileName)
	_, err := os.Stat(fmt.Sprintf("%s/%s", FilePath, fileName))
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "File not found.", http.StatusNotFound)
			return
		}
	}

	fmt.Fprintf(w, "Downloading content of file from: %s/%s\n", FilePath, fileName)
	fmt.Fprintf(w, "Content:\n")
	http.ServeFile(w, req, fmt.Sprintf("%s/%s", FilePath, fileName))
	w.WriteHeader(http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Delete file...")
	path := req.URL.Path
	e := strings.Split(path, "/")
	fileName := e[len(e)-1]
	fmt.Fprintf(w, "Deleting file %s\n", fileName)
	completeFileName := fmt.Sprintf("%s/%s", FilePath, fileName)
	_, err := os.Stat(completeFileName)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "File not found.", http.StatusNotFound)
			return
		}
	}

	err = os.Remove(completeFileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Successfully deleted the file %s\n", fileName)
	w.WriteHeader(http.StatusAccepted)
}
