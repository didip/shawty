package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/didip/shawty/handlers"
	"github.com/didip/shawty/storages"
	"github.com/mitchellh/go-homedir"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	dir, _ := homedir.Dir()
	storage := &storages.Filesystem{}
	err := storage.Init(filepath.Join(dir, "shawty"))
	if err != nil {
		log.Fatalln(err.Error())
	}

	http.Handle("/", handlers.EncodeHandler(storage))
	http.Handle("/dec/", handlers.DecodeHandler(storage))
	http.Handle("/red/", handlers.RedirectHandler(storage))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
