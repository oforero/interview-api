package topics

import (
	"fmt"
	"testing"
)

func TestConstructTopic(t *testing.T) {
	msg := "Test Topic"
	tp := NewTopic(msg)
	t.Logf("Got new message with id %v", tp.id)
	if msg != tp.msg {
		t.Errorf("Topic construction error: got %v want %v", tp.msg, msg)
	}
}

func TestGetKnownTopic(t *testing.T) {
	msg := "Test Topic"
	tp := NewTopic(msg)
	t.Logf("Got new message with id %v", tp.id)
	tp2, err := GetTopic(tp.id)
	if err != nil {
		t.Error(err)
	}
	if tp != tp2 {
		t.Errorf("Topic retrieval error: got %v want %v", tp2, tp)
	}
}

func TestUniqueIds(t *testing.T) {
	ids := make(map[string]int)
	for i := 0; i < 10000; i++ {
		msg := fmt.Sprintf("Topic %i", i)
		tp := NewTopic(msg)
		_, ok := ids[tp.id]
		if ok {
			t.Errorf("Repeated unique ID: %v", tp.id)
		}
	}
}

func TestUpvoteTopic(t *testing.T) {
	msg := "Test Topic"
	tp := NewTopic(msg)
	t.Logf("Got new message with id %v", tp.id)
	for i := 1; i < 10; i++ {
		Upvote(tp.id)
		tp, _ = GetTopic(tp.id)
		if tp.upvotes != i {
			t.Errorf("Topic upvote error: got %v want %v", tp.upvotes, i)
		}
	}
}

func TestJsonEncoding(t *testing.T) {
	msg := "Test Topic"
	tp := NewTopic(msg)
	t.Logf("Got new message with id %v", tp.id)
	tjson := EncodeToJSON(tp)
	tstr := fmt.Sprintf("{\"downvotes\":0,\"id\":\"%v\",\"msg\":\"%v\",\"upvotes\":0}", tp.id, tp.msg)
	if tjson != tstr {
		t.Errorf("Bad encoding!\n  got: %v\n want: %v", tjson, tstr)
	}
	t.Logf("Topic as json: %v", tjson)
}
