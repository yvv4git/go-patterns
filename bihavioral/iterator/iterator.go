package iterator

// Iterator - used for iterate elements of collection
type Iterator interface {
	Next() bool
	Get() interface{}
}
