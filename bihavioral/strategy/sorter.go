package strategy

// Sorter - common contract for all sorting algorithms
type Sorter interface {
	Sort([]int) []int
}
