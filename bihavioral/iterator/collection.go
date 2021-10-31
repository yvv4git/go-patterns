package iterator

// Collection - contract for iterable structure
type Collection interface {
	GetIterator() Iterator
}
