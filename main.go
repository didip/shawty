package main

import (
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

type Base36Url struct {
	Root string
}

func (s *Base36Url) Init(root string) {
	s.Root = root
	os.MkdirAll(s.Root, 0744)
}

func (s *Base36Url) Save(url string) string {
	files, _ := ioutil.ReadDir(s.Root)
	code := strconv.FormatUint(uint64(len(files)+1), 36)

	ioutil.WriteFile(filepath.Join(s.Root, code), []byte(url), 0744)
	return code
}

func (s *Base36Url) Load(code string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(s.Root, code))
}

func (s *Base36Url) EncodeHandler(w http.ResponseWriter, r *http.Request) {
	url := r.PostFormValue("url")
	if url != "" {
		w.Write([]byte(s.Save(url)))
	}
}

func (s *Base36Url) DecodeHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[len("/dec/"):]
	url, err := s.Load(code)

	if err == nil {
		w.Write(url)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: URL Not Found"))
	}
}

func (s *Base36Url) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[len("/red/"):]
	url, err := s.Load(code)

	if err == nil {
		http.Redirect(w, r, string(url), 301)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("URL Not Found"))
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	dir, _ := homedir.Dir()
	storage := &Base36Url{}
	storage.Init(filepath.Join(dir, "shawty"))

	http.HandleFunc("/", storage.EncodeHandler)
	http.HandleFunc("/dec/", storage.DecodeHandler)
	http.HandleFunc("/red/", storage.RedirectHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
