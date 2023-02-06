package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChocolateBall(t *testing.T) {
	emptyCup := new(EmptyCup)
	chocolateBall := NewChocolateBall(emptyCup) // Use decorator.
	assert.Equal(t, float32(60), chocolateBall.GetPrice())
}
