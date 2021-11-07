package observer

// Event - used as some event
type Event struct {
	message string
}

// NewEvent - simple factory for create instance of Event
func NewEvent(message string) Event {
	return Event{
		message: message,
	}
}

// Message - used for get message
func (e *Event) Message() string {
	return e.message
}
