package main

import (
	"fmt"
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

	tpl, err := ace.Load("templates/main", "", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, topics); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func newTopicHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	fmt.Printf("Got message: %v", msg)
	apierr := client.NewTopic(msg)
	if apierr != nil {
		http.Error(w, apierr.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 301)

}

func main() {
	client = web.NewDiggService(nil)
	http.HandleFunc("/", handler)
	http.HandleFunc("/newtopic", newTopicHandler)
	http.ListenAndServe(":8080", nil)
}
