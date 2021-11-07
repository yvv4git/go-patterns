package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrategy(t *testing.T) {
	data := []int{2, 1, 3, 5, 4, 8, 7, 6, 9}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	strategyBubble := NewStrategy(new(Bubble), data)
	assert.Equal(t, want, strategyBubble.Sort())

	strategyComb := NewStrategy(new(Comb), data)
	assert.Equal(t, want, strategyComb.Sort())
}
