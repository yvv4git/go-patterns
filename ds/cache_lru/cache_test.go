package cache_lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCacheLRU(t *testing.T) {
	cache := New(5)
	cache.Put("1", "first")
	cache.Put("2", "second")
	cache.Put("3", "third")
	cache.Put("4", "fourth")
	cache.Put("5", "fifth")

	valFirst := cache.Get("1")
	assert.Equal(t, "first", valFirst)

	valFirst = cache.Get("1")
	assert.Equal(t, "first", valFirst)

	valSecond := cache.Get("2")
	assert.Equal(t, "second", valSecond)

	// This value will fall into the place of the third element.
	cache.Put("6", "sixth")

	valThird := cache.Get("3")
	assert.Nil(t, valThird)
}
