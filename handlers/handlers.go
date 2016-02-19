// Package handlers provides HTTP request handlers.
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thomaso-mirodin/shawty/storage"
)

func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, mainhtml)
}

func GetShortHandler(store storage.Storage) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		code := p.ByName("short")

		url, err := store.Load(code)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "URL Not Found.")
			return
		}

		switch r.Header.Get("Accept") {
		case "application/json":
			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			err := json.NewEncoder(w).Encode(map[string]string{"short": code, "url": url})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "text/plain":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			fmt.Fprintln(w, url)
		default:
			http.Redirect(w, r, url, http.StatusFound)
		}
	}
}

func SetShortHandler(store storage.Storage) httprouter.Handle {
	named, namedOk := store.(storage.NamedStorage)
	unnamed, unnamedOk := store.(storage.UnnamedStorage)

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Get URL from post params based on Content-Type
		var (
			url string
			err error
		)
		switch r.Header.Get("Content-Type") {
		case "application/json":
			http.Error(w, "{'error': 'Content-Type of application/json not yet supported'}", http.StatusNotImplemented)
		default:
			url = r.PostFormValue("url")
			if url == "" {
				http.Error(w, "No URL Provided", http.StatusBadRequest)
				return
			}
		}

		// Save URL to the storage layer and get the final short code
		code := p.ByName("short")
		if code == "" {
			if !unnamedOk {
				http.Error(w, "Current storage layer does not support storing an unnamed url", http.StatusBadRequest)
				return
			}

			code, err = unnamed.Save(url)
		} else {
			if !namedOk {
				http.Error(w, "Current storage layer does not support storing a named url", http.StatusBadRequest)
				return
			}

			err = named.SaveName(code, url)
		}
		if err != nil {
			http.Error(w, fmt.Sprintf("Storage layer failed to save link: %s", err), http.StatusInternalServerError)
		}

		// Return the short code formatted based on Accept headers
		switch r.Header.Get("Accept") {
		case "application/json":
			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			err := json.NewEncoder(w).Encode(map[string]string{"short": code, "url": url})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "text/plain":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			fmt.Fprintln(w, code)
		default:
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			fmt.Fprintln(w, code)
		}
	}

}
