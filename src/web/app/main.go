package main

import (
	"fmt"
	"github.com/yosssi/ace"
	"log"
	"net/http"
	"time"
	"web"
)

var client *web.DiggService

type DiggHandler func(w http.ResponseWriter, r *http.Request)

func logging(name string, handler DiggHandler) DiggHandler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t_start := time.Now()
		handler(w, r)
		t_end := time.Now()
		log.Printf("[%s] [%s] %q %v\n", name, r.Method, r.URL.String(), t_end.Sub(t_start))
	}

	return fn
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
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
	apierr := client.NewTopic(msg)
	if apierr != nil {
		http.Error(w, apierr.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 303)
}

func upvoteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Printf("Got upvote request for: %v\n", id)
	apierr := client.Upvote(id)
	if apierr != nil {
		http.Error(w, apierr.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 303)
}

func downvoteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Printf("Got upvote request for: %v\n", id)
	apierr := client.Downvote(id)
	if apierr != nil {
		http.Error(w, apierr.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 303)
}

func main() {
	client = web.NewDiggService(nil)
	http.HandleFunc("/newtopic", logging("newtopic", newTopicHandler))
	http.HandleFunc("/upvote", logging("upvote", upvoteHandler))
	http.HandleFunc("/downvote", logging("downvote", downvoteHandler))
	http.HandleFunc("/", logging("/", rootHandler))
	http.ListenAndServe(":8080", nil)
}
