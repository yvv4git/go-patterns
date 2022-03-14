package singleton

import "sync"

var once sync.Once

type singleton map[string]string

var (
	version singleton
)

// NewSingleton - used to provide an instance of an entity
func NewSingleton() singleton {

	once.Do(func() {
		version = make(singleton)
	})

	return version
}
