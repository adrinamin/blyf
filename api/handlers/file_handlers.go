package handlers

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "path/filepath"
    "os"
    "log"
)

const (
	FilePath = "files"
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


