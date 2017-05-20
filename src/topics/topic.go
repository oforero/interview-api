package topics

type topic struct {
	id        int
	msg       string
	upvotes   int
	downvotes int
}

func NewTopic(msg string) *topic {
	return &topic{0, msg, 0, 0}
}
