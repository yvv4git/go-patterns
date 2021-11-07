package strategy

// Bubble - used as concrete sort realisation
type Bubble struct{}

// Sort - bubble sort algorithm
func (b *Bubble) Sort(data []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(data)-1; i++ {
			if data[i+1] < data[i] {
				Swap(data, i, i+1)
				swapped = true
			}
		}
	}

	return data
}
