package main

import (
	"log"
	"net/http"

	"image_upload/internal/upload"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/upload", upload.Handler)
	mux.HandleFunc("/list", upload.ListHandler)
	mux.HandleFunc("/delete", upload.DeleteHandler)
	log.Println("Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
