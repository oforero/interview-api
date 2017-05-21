package topics

import (
	"encoding/json"
	"github.com/ventu-io/go-shortid"
)

type topic struct {
	Id        string
	Msg       string
	Upvotes   int
	Downvotes int
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

func EncodeToJSON(t topic) string {
	tjson, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(tjson)
}
