package strategy

// Comb - used as concrete sort realisation
type Comb struct{}

// Sort - comb sort algorithm
func (c *Comb) Sort(data []int) []int {
	if len(data) < 2 {
		return data
	}

	for gap := len(data); ; {
		if gap > 1 {
			gap = gap * 4 / 5
		}
		swapped := false
		for i := 0; ; {
			if data[i] > data[i+gap] {
				data[i], data[i+gap] = data[i+gap], data[i]
				swapped = true
			}
			i++
			if i+gap >= len(data) {
				break
			}
		}
		if gap == 1 && !swapped {
			break
		}
	}

	return data
}
