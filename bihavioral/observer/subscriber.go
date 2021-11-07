package observer

// Subscriber - client who wants to receive notifications
type Subscriber struct {
	name    string
	mailBox []string
}

// NewSubscriber - simple factory for create instance of Subscriber
func NewSubscriber(name string) *Subscriber {
	return &Subscriber{
		name:    name,
		mailBox: make([]string, 0),
	}
}

// Receive - receive message from publisher
func (s *Subscriber) Receive(event Event) {
	s.mailBox = append(s.mailBox, event.Message())
}

// Get - used for getting message
func (s *Subscriber) Get(messageIndex int) string {
	if len(s.mailBox) > messageIndex {
		return s.mailBox[messageIndex]
	}

	return ""
}
