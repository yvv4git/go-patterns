package memento

// Editor - used as simple text editor
type Editor struct {
	text string
}

// Type - add string
func (e *Editor) Type(text string) {
	e.text = e.text + "\n" + text
}

// Save - used for save state
func (e *Editor) Save() Memento {
	return NewMementoEditor(e.text)
}

// Restore - used for restore state from memento
func (e *Editor) Restore(memento Memento) {
	e.text = memento.Restore()
}
