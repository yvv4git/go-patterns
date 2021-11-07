package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmailPublisher(t *testing.T) {
	const message = "Free seminar on Wednesday at 5 pm"

	emailPublisher := NewEmail()

	subscriber1 := NewSubscriber("Vladimir")
	emailPublisher.Subscribe(subscriber1)

	subscriber2 := NewSubscriber("Anonymous")
	emailPublisher.Subscribe(subscriber2)

	event := NewEvent(message)
	emailPublisher.Notify(event)

	assert.Equal(t, message, subscriber1.Get(0))
	assert.Equal(t, message, subscriber2.Get(0))
}
