package topics

import (
	"github.com/ventu-io/go-shortid"
)

type topic struct {
	id        string
	msg       string
	upvotes   int
	downvotes int
}

func init() {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		panic(err)
	}
	shortid.SetDefault(sid)
}

func NewTopic(msg string) *topic {
	id, err := shortid.Generate()
	if err != nil {
		panic(err)
	}
	return &topic{id, msg, 0, 0}
}
