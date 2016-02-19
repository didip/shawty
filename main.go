package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/thomaso-mirodin/shawty/handlers"
	"github.com/thomaso-mirodin/shawty/storage"
)

func main() {
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// path := filepath.Join(dir, "shawty", "filesystem_db")
	// store, err := storage.NewFilesystem(path)
	// if err := err {
	// 	log.Fatalf("Failed to create filesystem '%s' because '%s'", path, err)
	// }
	store, err := storage.NewInmem()
	if err != nil {
		log.Fatalf("Failed to create inmem storage because '%s'", err)

	}

	n := negroni.Classic()

	mux := httprouter.New()
	mux.GET("/", handlers.Index)

	mux.GET("/:short", handlers.GetShortHandler(store))
	mux.HEAD("/:short", handlers.GetShortHandler(store))

	mux.POST("/", handlers.SetShortHandler(store))
	mux.PUT("/", handlers.SetShortHandler(store))
	mux.POST("/:short", handlers.SetShortHandler(store))
	mux.PUT("/:short", handlers.SetShortHandler(store))

	n.UseHandler(mux)

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = http.ListenAndServe(net.JoinHostPort(host, port), n)
	if err != nil {
		log.Fatal(err)
	}
}
