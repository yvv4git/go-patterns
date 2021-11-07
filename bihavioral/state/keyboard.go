package state

// Keyboard - object that has different states
type Keyboard struct {
	state State
}

// NewKeyboard - simple factory for create instance of Keyboard
func NewKeyboard(state State) *Keyboard {
	return &Keyboard{
		state: state,
	}
}

// SetState - used for set state of keyboard
func (k *Keyboard) SetState(state State) {
	k.state = state
}

// Type - we are working in one of the possible states
func (k *Keyboard) Type(words string) string {
	return k.state.Write(words)
}
