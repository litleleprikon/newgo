package server

import (
	"fmt"
	"io"

	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/hello/name", getHelloName)
	return mux
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getHelloName(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	fmt.Printf("got /hello/name request\n")
	fmt.Fprintf(w, "Hello %s\n", name)
}
