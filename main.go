package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strconv"
)

type Base36Url struct {
	Root string
}

func (s *Base36Url) Init(root string) {
	s.Root = root
	os.MkdirAll(s.Root, 0744)
}

func (s *Base36Url) Path(code string) string {
	return s.Root + code
}

func (s *Base36Url) Save(url string) string {
	files, _ := ioutil.ReadDir(s.Path(""))
	code := strconv.FormatUint(uint64(len(files)+1), 36)

	ioutil.WriteFile(s.Path(code), []byte(url), 0744)
	return code
}

func (s *Base36Url) EncodeHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url != "" {
		fmt.Fprintf(w, s.Save(url))
	}
}

func (s *Base36Url) DecodeHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	code := params["code"]
	url, err := ioutil.ReadFile(s.Path(code))

	if err == nil {
		fmt.Fprintf(w, string(url))
	} else {
		fmt.Fprintf(w, "Error: %v", err)
	}
}

func (s *Base36Url) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := mux.Vars(r)["code"]
	url, err := ioutil.ReadFile(s.Path(code))

	if err == nil {
		http.Redirect(w, r, string(url), 301)
	}
}

func main() {
	usr, _ := user.Current()
	storage := &Base36Url{}
	storage.Init(usr.HomeDir + "/shawty/")

	r := mux.NewRouter()
	r.HandleFunc("/", storage.EncodeHandler).Methods("POST")
	r.HandleFunc("/dec/{code}", storage.DecodeHandler).Methods("GET")
	r.HandleFunc("/red/{code}", storage.RedirectHandler).Methods("GET")
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
