package topics

import (
	"encoding/json"
	"fmt"
	"github.com/ventu-io/go-shortid"
)

type topic struct {
	Id        string
	Msg       string
	Upvotes   int
	Downvotes int
}

var index map[string]int
var db []*topic

func init() {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		panic(err)
	}
	shortid.SetDefault(sid)

	db = make([]*topic, 0, 20)
	index = make(map[string]int)
}

func NewTopic(msg string) *topic {
	id, err := shortid.Generate()
	if err != nil {
		panic(err)
	}
	t := &topic{id, msg, 0, 0}
	db = append(db, t)
	index[t.Id] = len(db) - 1
	return t
}

func GetTopic(id string) *topic {
	ix, ok := index[id]
	if !ok {
		panic(fmt.Sprintf("Topic with id (%v) not found", id))
	}
	return db[ix]
}

func EncodeToJSON(t topic) string {
	tjson, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(tjson)
}
