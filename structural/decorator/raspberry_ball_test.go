package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRaspberryBall(t *testing.T) {
	emptyCup := new(EmptyCup)
	raspberryBall := NewRaspberryBall(emptyCup) // Use decorator.
	assert.Equal(t, float32(70.0), raspberryBall.GetPrice())
}
