package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Uploading Files!\n")

	// Parse the multipart form with a maximum size of 10 MB
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Fprintf(w, "Error parsing form: %v", err)
		wg.Done()
		return
	}

	// Retrieve the file from posted form-data
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Fprintf(w, "Error retrieving the file: %v", err)
		wg.Done()
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
		wg.Done()
		return
	}

	// Write temporary file on our server
	tempFile, err := os.CreateTemp(tempDir, "upload-*.png")
	if err != nil {
		fmt.Fprintf(w, "Error creating temporary file: %v", err)
		wg.Done()
		return
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "Error reading file: %v", err)
		wg.Done()
		return
	}
	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Successfully uploaded!\n")
	wg.Done()
}

func setupRoutes() *http.Server {
	http.HandleFunc("/upload", uploadFile)
	server := &http.Server{Addr: ":8080"}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe(): %v\n", err)
		}
	}()
	return server
}

func main() {
	fmt.Println("Starting server on :8080")
	wg.Add(1)
	server := setupRoutes()

	wg.Wait()
	fmt.Println("Shutting down the server...")
	if err := server.Close(); err != nil {
		fmt.Printf("Error shutting down the server: %v\n", err)
	}
}
