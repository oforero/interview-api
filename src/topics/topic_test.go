package topics

import "testing"

func TestConstructTopic(t *testing.T) {
	msg := "Test Topic"
	tp := NewTopic(msg)
	if msg != tp.msg {
		t.Errorf("Topic construction error: got %v want %v", tp.msg, msg)
	}
}
