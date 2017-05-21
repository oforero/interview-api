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
