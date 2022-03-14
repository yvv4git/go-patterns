package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingletone(t *testing.T) {
	singleton := NewSingleton()
	singleton["key-1"] = "value-1"

	singleton2 := NewSingleton()
	assert.Equal(t, "value-1", singleton2["key-1"])
}
