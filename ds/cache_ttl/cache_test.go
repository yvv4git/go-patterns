package cache_ttl

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	const (
		key = "some-key"
		val = "some-val"
	)

	cache := NewCache(time.Minute, 2*time.Minute)

	cache.Set(key, val, 2*time.Minute)

	item, ok := cache.Get(key)
	assert.True(t, ok)
	assert.Equal(t, val, item)
}
