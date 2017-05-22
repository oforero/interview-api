package web

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

const baseURL = "http://api:8000/"

type topic struct {
	ID        string `json:"id"`
	Msg       string `json:"msg"`
	Upvotes   int    `json:"upvotes"`
	Downvotes int    `json:"downvotes"`
}

type DiggService struct {
	sling *sling.Sling
}

func NewDiggService(httpClient *http.Client) *DiggService {
	return &DiggService{
		sling: sling.New().Client(httpClient).Base(baseURL),
	}
}

func (s *DiggService) GetTopics() ([]topic, error) {
	topics := new([]topic)
	_, err := s.sling.New().Get("topics").ReceiveSuccess(topics)
	if err == nil {
		fmt.Print("Error getting topics")
	}
	return *topics, err
}
