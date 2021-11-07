package mediator

type Messenger interface {
	Receive(message string)
}
