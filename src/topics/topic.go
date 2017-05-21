package topics

import (
	"encoding/json"
	"errors"
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

func getTopic(id string) (*topic, int, error) {
	ix, ok := index[id]
	if !ok {
		return nil, 0, errors.New(fmt.Sprintf("Topic with id (%v) not found", id))
	}
	return db[ix], ix, nil
}

func GetTopic(id string) (*topic, error) {
	tp, _, err := getTopic(id)
	return tp, err
}

func score(ix int) int {
	tp := db[ix]
	return tp.Upvotes - tp.Downvotes
}

func swap(ix1 int, ix2 int) {
	tp1 := db[ix1]
	tp2 := db[ix2]
	db[ix1] = tp2
	index[tp2.Id] = ix1
	db[ix2] = tp1
	index[tp1.Id] = ix2
}

func compareUpAndSwap(ix int) {
	if ix == 0 {
		return
	}
	ix_prev := ix - 1
	if score(ix_prev) < score(ix) {
		swap(ix_prev, ix)
	}
}

func Upvote(id string) error {
	tp, ix, err := getTopic(id)
	if err != nil {
		return err
	}
	tp.Upvotes = tp.Upvotes + 1
	compareUpAndSwap(ix)
	return nil
}

func EncodeToJSON(t topic) string {
	tjson, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(tjson)
}
