package observer

// Publisher - contract for all publishers
type Publisher interface {
	Notify(event Event)
}
