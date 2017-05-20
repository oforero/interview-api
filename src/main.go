package main

import (
	"net/http"
)

type helloHandler struct{}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello again Go!"))
}

func main() {

	http.ListenAndServe(":8000", helloHandler{})
}
