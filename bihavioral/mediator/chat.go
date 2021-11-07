package mediator

import "sync"

// Chat - some chat
type Chat struct {
	sync.RWMutex
	messages []string
}

func NewChat() *Chat {
	return &Chat{
		messages: make([]string, 0),
	}
}

// Receive - used for receive message for chat
func (s *Chat) Receive(message string) {
	s.RLock()
	defer s.RUnlock()
	s.messages = append(s.messages, message)
}

// Messages - get all messages
func (s *Chat) Messages() []string {
	s.RLock()
	defer s.RUnlock()
	return s.messages
}
