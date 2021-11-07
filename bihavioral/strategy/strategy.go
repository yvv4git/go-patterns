package strategy

// Strategy - context of strategy
type Strategy struct {
	algorithm Sorter
	data      []int
}

// NewStrategy - simple factory for create instance of Strategy
func NewStrategy(algorithm Sorter, data []int) *Strategy {
	return &Strategy{
		algorithm: algorithm,
		data:      data,
	}
}

// Sort - used for sort elements with concrete algorithm
func (s *Strategy) Sort() []int {
	return s.algorithm.Sort(s.data)
}
