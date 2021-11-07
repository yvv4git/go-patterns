package strategy

// Swap is used for swapping two values in slice.
func Swap(data []int, i, j int) {
	tmp := data[j]
	data[j] = data[i]
	data[i] = tmp
}
