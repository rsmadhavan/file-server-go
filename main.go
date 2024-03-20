package main

import ( 
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)


const uploadDir = "./uploads"

func main() {
    // Create uploads directory if it doesn't exist
    if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
        os.Mkdir(uploadDir, os.ModePerm)
    }

    http.HandleFunc("/upload", uploadHandler)
    http.HandleFunc("/files/", fileHandler)

    fmt.Println("Server listening on localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Save the file to disk
    filepath := filepath.Join(uploadDir, handler.Filename)
    f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }
    defer f.Close()

    _, err = io.Copy(f, file)
    if err != nil {
        http.Error(w, "Error saving file", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
    filepath := filepath.Join(uploadDir, r.URL.Path[len("/files/"):])
    http.ServeFile(w, r, filepath)
}