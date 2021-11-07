package mediator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStdoutMessenger(t *testing.T) {
	chat := NewChat()
	client1 := NewClient(chat)
	client2 := NewClient(chat)

	client1.Send("Hi, how are you?")
	client2.Send("Hi! Everything is great.")

	assert.Equal(t, 2, len(chat.Messages()))
}
