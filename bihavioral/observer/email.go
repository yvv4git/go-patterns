package observer

import "sync"

// Email - used as concrete publisher for email notification
type Email struct {
	sync.RWMutex
	subscribers []*Subscriber
}

// NewEmail - simple factory for create instance of Email
func NewEmail() *Email {
	return &Email{
		subscribers: make([]*Subscriber, 0),
	}
}

// Subscribe - add new subscriber
func (e *Email) Subscribe(subscriber *Subscriber) {
	e.RLock()
	defer e.RUnlock()
	e.subscribers = append(e.subscribers, subscriber)
}

// Notify - used for send event for all subscribers
func (e *Email) Notify(event Event) {
	e.RLock()
	defer e.RUnlock()

	for _, subscriber := range e.subscribers {
		subscriber.Receive(event)
	}
}
