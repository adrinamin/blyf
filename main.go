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
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	file, handler, err := req.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// todo: Create a destination file
	// todo: Copy the uploaded file's content to the destination file

	fmt.Fprintf(w, "Upload of %s was successful.", handler.Filename)

}

func downloadHandler(w http.ResponseWriter, req *http.Request) {
	fileDownloadMessage := "File download"
	fmt.Println("Download file...")
	io.WriteString(w, fmt.Sprintf("%s\n", fileDownloadMessage))
}

func deleteHandler(w http.ResponseWriter, req *http.Request) {
	fileDeleteMessage := "File delete"
	fmt.Println("Delete file...")
	io.WriteString(w, fmt.Sprintf("%s\n", fileDeleteMessage))
}
