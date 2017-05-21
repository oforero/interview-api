package main

import (
	"net/http"
	"topics"
)

type helloHandler struct{}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("Hello world"))
}

func GetTopicsHandler(w http.ResponseWriter, r *http.Request) {
	topics, _ := topics.GetTopicsAsJSON()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(topics)
}

func NewTopicHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	topics.NewTopic(msg)
	tps, _ := topics.GetTopicsAsJSON()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(tps)
}

func UpvoteTopicHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	topics.Upvote(id)
	tps, _ := topics.GetTopicsAsJSON()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(tps)
}

func DownvoteTopicHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	topics.Downvote(id)
	tps, _ := topics.GetTopicsAsJSON()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(tps)
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/topics", GetTopicsHandler)
	http.HandleFunc("/topics/new", NewTopicHandler)
	http.HandleFunc("/topics/upvote", UpvoteTopicHandler)
	http.HandleFunc("/topics/downvote", DownvoteTopicHandler)
	http.ListenAndServe(":8000", nil)
}
