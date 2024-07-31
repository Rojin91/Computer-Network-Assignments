package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading Files!\n")

	// 1. Parse the multipart form with a maximum size of 10 MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Fprintf(w, "Error parsing form: %v", err)
		return
	}

	// 2. Retrieve the file from posted form-data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Fprintf(w, "Error retrieving the file: %v", err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Ensure the directory exists
	const tempDir = "temp-images"
	if err := os.MkdirAll(tempDir, os.ModePerm); err != nil {
		fmt.Fprintf(w, "Error creating directory: %v", err)
		return
	}

	// 3. Write temporary file on our server
	tempFile, err := os.CreateTemp(tempDir, "upload-*.png")
	if err != nil {
		fmt.Fprintf(w, "Error creating temporary file: %v", err)
		return
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "Error reading file: %v", err)
		return
	}
	tempFile.Write(fileBytes)

	// 4. Return whether or not this has been successful
	fmt.Fprintf(w, "Successfully uploaded!\n")
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":5500", nil)
}

func main() {
	fmt.Println("Uploading Files!")
	setupRoutes()
}
