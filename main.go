package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-homedir"
	"github.com/thomaso-mirodin/shawty/handlers"
	"github.com/thomaso-mirodin/shawty/storages"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	dir, _ := homedir.Dir()
	storage := &storages.Filesystem{}
	err := storage.Init(filepath.Join(dir, "shawty"))
	if err != nil {
		log.Fatal(err)
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
