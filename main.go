package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/thomaso-mirodin/shawty/handlers"
	"github.com/thomaso-mirodin/shawty/storage"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(dir, "shawty", "filesystem_db")
	store, err := storage.NewFilesystem(path)
	if err != nil {
		log.Fatalf("Failed to create filesystem '%s' because '%s'", path, err)
	}

	r := httprouter.New()
	r.GET("/", handlers.Index)
	r.GET("/:short", handlers.RedirectHandler(store))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
