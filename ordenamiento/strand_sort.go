package ordenamiento

// Creado por ChatGPT
// Adaptado para usar punteros
func StrandSort(arr *[]int) []int {
	if len(*arr) <= 1 {
		return *arr
	}

	// Split the array into sorted and unsorted parts
	sorted := []int{(*arr)[0]}
	*arr = (*arr)[1:]

	for len(*arr) > 0 {
		// Split the remaining array into subarrays that are sorted
		sub := []int{(*arr)[0]}
		*arr = (*arr)[1:]

		for i := 0; i < len(sorted); i++ {
			if sorted[i] < sub[0] {
				continue
			}
			j := i
			for j < len(sorted) && sub[len(sub)-1] > sorted[j] {
				sub = append(sub, sorted[j])
				j++
			}
			sorted = append(sorted[:i], append(sub, sorted[i:]...)...)
			break
		}

		if len(sorted) == 0 || sub[len(sub)-1] > sorted[len(sorted)-1] {
			sorted = append(sorted, sub...)
		}
	}

	return sorted
}
