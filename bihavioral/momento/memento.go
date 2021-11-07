package momento

// Memento - contract for all mementos
type Memento interface {
	Restore() string
}
