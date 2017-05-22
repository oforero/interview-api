package main

import (
	"github.com/yosssi/ace"
	"net/http"
	"web"
)

var client *web.DiggService

func handler(w http.ResponseWriter, r *http.Request) {
	topics, apierr := client.GetTopics()
	if apierr != nil {
		http.Error(w, apierr.Error(), http.StatusInternalServerError)
		return
	}

	tpl, err := ace.Load("main", "", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, topics); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	client = web.NewDiggService(nil)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
