package memento

// MementoEditor - some realisation of Memento
type MementoEditor struct {
	state string
}

// NewMementoEditor - simple factory for create instance of MementoEditor
func NewMementoEditor(text string) MementoEditor {
	return MementoEditor{
		state: text,
	}
}

// Restore - restore state
func (m MementoEditor) Restore() string {
	return m.state
}
