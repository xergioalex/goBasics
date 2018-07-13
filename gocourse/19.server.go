package main

import (
	"net/http"
	"io"
)

func handler (w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hola Servidor Go!")
}

func main () {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ii", func (w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Route: ii!")
	})
	http.ListenAndServe(":8000", nil)
}