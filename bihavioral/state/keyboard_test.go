package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyboardState(t *testing.T) {
	stateLowerCase := new(LowerCase)
	keyboard := NewKeyboard(stateLowerCase)
	assert.Equal(t, "my text", keyboard.Type("My text"))

	stateUpperCase := new(UpperCase)
	keyboard.SetState(stateUpperCase)
	assert.Equal(t, "MY TEXT", keyboard.Type("My text"))
}
