package factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorySimple(t *testing.T) {
	person := NewPerson("Vladimir", 32)

	assert.Equal(t, "Vladimir", person.Name)
	assert.Equal(t, 32, person.Age)
	assert.Equal(t, "I'am Vladimir and i'am 32", person.GetName())
}
