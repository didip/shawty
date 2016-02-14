// Package handlers provides HTTP request handlers.
package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/thomaso-mirodin/shawty/storage"
)

func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Hello world!")
}

func RedirectHandler(storage storage.Storage) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		code := ps.ByName("short")

		url, err := storage.Load(code)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "URL Not Found.")
			return
		}

		http.Redirect(w, r, string(url), http.StatusFound)
	}
}

// func UnknownRedirectHandler(storage storages.Unnamed) httprouter.Handler {
// 	handleFunc := func(w http.ResponseWriter, r *http.Request) {
// 		if url := r.PostFormValue("url"); url != "" {
// 			code, err := storage.Save(url)
// 			if err != nil {
// 				http.Error(w, err, 500)
// 				return
// 			}
// 			w.Write([]byte(code))
// 		}
// 	}
// }

// 	return http.HandlerFunc(handleFunc)
// }

// func DecodeHandler(storage storages.Storage) http.Handler {
// 	handleFunc := func(w http.ResponseWriter, r *http.Request) {
// 		code := r.URL.Path[len("/dec/"):]

// 		url, err := storage.Load(code)
// 		if err != nil {
// 			w.WriteHeader(http.StatusNotFound)
// 			w.Write([]byte("URL Not Found. Error: " + err.Error() + "\n"))
// 			return
// 		}

// 		w.Write([]byte(url))
// 	}

// 	return http.HandlerFunc(handleFunc)
// }
