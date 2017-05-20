package main

import (
	"net/http"
)

type helloHandler struct{}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("Hello again Go!"))
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8000", nil)
}
