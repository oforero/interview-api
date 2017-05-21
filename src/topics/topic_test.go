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
