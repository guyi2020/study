package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("PATH: ", r.URL.Path)
	fmt.Println("SCHEME: ", r.URL.Scheme)
	fmt.Println("METHOD: ", r.Method)
	fmt.Println()
	fmt.Fprintf(w, "<h1>Index Page</h1>")
	io.WriteString(w, "this is a writer! ")
}

func main() {
	http.HandleFunc("/", HandleIndex)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}
