package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Specify the subdirectory name
	subdir := "tmp"
	// Create the full path to the subdirectory
	tmpDir := filepath.Join(currentDir, subdir)

	// Create the subdirectory if it doesn't exist
	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		fmt.Println("Error creating subdirectory:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `file`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile(tmpDir, "upload-*.png")
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	defer tempFile.Close()

	basePath := "http://192.168.112.26:3001" + "/public/" + filepath.Base(tempFile.Name())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fileResponse{
		ImageUrl: basePath,
	})
}

type fileResponse struct {
	ImageUrl string `json:"imageUrl"`
}
