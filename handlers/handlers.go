// Package handlers provides HTTP request handlers.
package handlers

import (
	"net/http"

	"github.com/didip/shawty/storages"
)

func EncodeHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if url := r.PostFormValue("url"); url != "" {
				w.Write([]byte(storage.Save(url)))
			}

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}

	return http.HandlerFunc(handleFunc)
}

func DecodeHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			code := r.URL.Path[len("/dec/"):]

			url, err := storage.Load(code)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("URL Not Found. Error: " + err.Error() + "\n"))
				return
			}

			w.Write([]byte(url))

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}

	return http.HandlerFunc(handleFunc)
}

func RedirectHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			code := r.URL.Path[len("/red/"):]

			url, err := storage.Load(code)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("URL Not Found. Error: " + err.Error() + "\n"))
				return
			}

			http.Redirect(w, r, string(url), 301)

		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}

	return http.HandlerFunc(handleFunc)
}
