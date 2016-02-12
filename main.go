package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/thomaso-mirodin/shawty/handlers"
	"github.com/thomaso-mirodin/shawty/storages"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(dir, "shawty")
	storage, err := storages.NewFilesystem(path)
	if err != nil {
		log.Fatalf("Failed to create filesystem '%s' because '%s'", path, err)
	}

	http.Handle("/", handlers.EncodeHandler(storage))
	http.Handle("/dec/", handlers.DecodeHandler(storage))
	http.Handle("/red/", handlers.RedirectHandler(storage))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
