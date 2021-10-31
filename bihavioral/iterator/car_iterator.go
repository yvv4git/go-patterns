package iterator

// CarIterator - used for work with CarCollection
type CarIterator struct {
	Cars  []*Car
	Index int
}

// Next - used to check for the presence of the following element
func (c *CarIterator) Next() bool {
	if c.Index < len(c.Cars) {
		return true
	}

	return false
}

// Get - used for get Car
func (c *CarIterator) Get() interface{} {
	if c.Next() {
		car := c.Cars[c.Index]
		c.Index++
		return car
	}

	return nil
}

// NewCarIterator - simple factory for create CarIterator instance
func NewCarIterator(list CarCollection) *CarIterator {
	return &CarIterator{
		Cars:  list,
		Index: 0,
	}
}
