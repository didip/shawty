package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"os/signal"
	"fmt"

	"github.com/didip/shawty/handlers"
	"github.com/didip/shawty/storages"
	"github.com/mitchellh/go-homedir"
)

func main() {

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

	// Graceful shutdown
	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, os.Kill)

	// Wait signal
	close := make(chan bool, 1)

	// Create a server
	server := &http.Server{Addr: fmt.Sprintf(":%s",  port)}

	// Start server
	go func() {
		log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			if err := server.ListenAndServe(); err != nil {
				if err != http.ErrServerClosed {
					log.Println(err.Error())
				} else {
					log.Println("Server Closed")
				}
				close <- true
			}
		}

	}()

	// Check for a closing signal
	go func() {
		sig := <-sigquit
		log.Printf("caught sig: %+v", sig)

		if err := server.Shutdown(nil); err != nil {
			log.Println("Unable to shut down server: " + err.Error())
			close <- true
		} else {
			log.Println("Server stopped")
			close <- true
		}
	}()

	<-close
}
