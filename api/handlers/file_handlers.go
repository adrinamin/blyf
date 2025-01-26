package handlers

import (
	"fmt"
	"github.com/adrinamin/blyf/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	FilePath    = "files"
	MaxFileSize = 10 * 1024 * 1024
)

func GetFilesHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Here are your current files:\n ")

	files, err := ioutil.ReadDir(FilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for _, file := range files {
		fmt.Fprintf(w, "%s\n", file.Name())
	}

	w.WriteHeader(http.StatusOK)
}

func UploadFileHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form, with a maximum memory of 32MB
	err := req.ParseMultipartForm(32 << 20)
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
		http.Error(w, "File upload failed because file name already exists.", http.StatusBadRequest)
		return
	}

	ext := filepath.Ext(completeFileName)
	if !utils.IsValidExtension(ext) {
		http.Error(w, "Invalid file extension", http.StatusBadRequest)
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

func DownloadFileHandler(w http.ResponseWriter, req *http.Request) {
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

func DeleteFileHandler(w http.ResponseWriter, req *http.Request) {
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
