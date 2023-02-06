package decorator

// EmptyCup - used as default empty cup. It's just an empty glass
type EmptyCup struct {
}

// GetPrice - used for return price.
func (e *EmptyCup) GetPrice() float32 {
	return 10
}
