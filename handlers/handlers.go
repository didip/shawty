// Package handlers provides HTTP request handlers.
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thomaso-mirodin/go-shorten/storage"
)

func getShortFromRequest(r *http.Request) (short string, err error) {
	short = r.URL.Path[1:]

	if short == "" {
		switch r.Header.Get("Content-Type") {
		case "application/json":
			err = fmt.Errorf("Content-Type of application/json not yet supported")
		default:
			short = r.PostFormValue("code")
		}
	}

	return
}

func getURLFromRequest(r *http.Request) (url string, err error) {
	switch r.Header.Get("Content-Type") {
	case "application/json":
		err = fmt.Errorf("Content-Type of application/json not yet supported")
		return
	default:
		url = r.PostFormValue("url")
	}

	if url == "" {
		err = fmt.Errorf("No URL Provided")
	}

	return
}

func GetShortHandler(store storage.Storage) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		short, err := getShortFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		url, err := store.Load(short)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "URL Not Found.")
			return
		}

		switch r.Header.Get("Accept") {
		case "application/json":
			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			err := json.NewEncoder(w).Encode(map[string]string{"short": short, "url": url})
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

	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		short, err := getShortFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		url, err := getURLFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if short == "" {
			if !unnamedOk {
				http.Error(w, "Current storage layer does not support storing an unnamed url", http.StatusBadRequest)
				return
			}

			short, err = unnamed.Save(url)
		} else {
			if !namedOk {
				http.Error(w, "Current storage layer does not support storing a named url", http.StatusBadRequest)
				return
			}

			err = named.SaveName(short, url)
		}
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to save '%s' to '%s' because: %s", url, short, err), http.StatusInternalServerError)
			return
		}

		// Return the short code formatted based on Accept headers
		switch r.Header.Get("Accept") {
		case "application/json":
			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			err := json.NewEncoder(w).Encode(map[string]string{"short": short, "url": url})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "text/html":
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintln(w, "<html>hello world</html")
		case "text/plain":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			fmt.Fprintln(w, short)
		default:
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			fmt.Fprintln(w, short)
		}
	}

}
