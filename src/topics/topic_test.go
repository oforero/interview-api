package topics

import (
	"fmt"
	"testing"
)

func TestConstructTopic(t *testing.T) {
	msg := "Test Topic"
	tp := NewTopic(msg)
	t.Logf("Got new message with id %v", tp.Id)
	if msg != tp.Msg {
		t.Errorf("Topic construction error: got %v want %v", tp.Msg, msg)
	}
}

func TestGetKnownTopic(t *testing.T) {
	msg := "Test Topic"
	tp := NewTopic(msg)
	t.Logf("Got new message with id %v", tp.Id)
	tp2, err := GetTopic(tp.Id)
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
		_, ok := ids[tp.Id]
		if ok {
			t.Errorf("Repeated unique ID: %v", tp.Id)
		}
	}
}

func TestUpvoteTopic(t *testing.T) {
	msg := "Test Topic"
	tp := NewTopic(msg)
	t.Logf("Got new message with id %v", tp.Id)
	for i := 1; i < 10; i++ {
		Upvote(tp.Id)
		tp, _ = GetTopic(tp.Id)
		if tp.Upvotes != i {
			t.Errorf("Topic upvote error: got %v want %v", tp.Upvotes, i)
		}
	}
}

func TestJsonEncoding(t *testing.T) {
	msg := "Test Topic"
	tp := NewTopic(msg)
	t.Logf("Got new message with id %v", tp.Id)
	tjson := EncodeToJSON(*tp)
	tstr := fmt.Sprintf("{\"Id\":\"%v\",\"Msg\":\"%v\",\"Upvotes\":0,\"Downvotes\":0}", tp.Id, tp.Msg)
	if tjson != tstr {
		t.Errorf("Bad encoding!\n  got: %v\n want: %v", tjson, tstr)
	}
	t.Logf("Topic as json: %v", tjson)
}
