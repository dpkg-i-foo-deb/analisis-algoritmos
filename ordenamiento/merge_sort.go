package ordenamiento

// Creado por BingAI
func MergeSort(src *[]int) {
	if len(*src) <= 1 {
		return
	}
	mid := len(*src) / 2
	left := make([]int, mid)
	right := make([]int, len(*src)-mid)
	copy(left, (*src)[:mid])
	copy(right, (*src)[mid:])
	MergeSort(&left)
	MergeSort(&right)
	merge(src, &left, &right)
}

// merge merges two sorted sub-arrays into one sorted array
func merge(dst *[]int, left *[]int, right *[]int) {
	i := 0 // index for left array
	j := 0 // index for right array
	k := 0 // index for merged array

	for i < len(*left) && j < len(*right) {
		if (*left)[i] <= (*right)[j] {
			(*dst)[k] = (*left)[i]
			i++
		} else {
			(*dst)[k] = (*right)[j]
			j++
		}
		k++
	}

	// copy remaining elements from left or right array if any
	for i < len(*left) {
		(*dst)[k] = (*left)[i]
		i++
		k++
	}

	for j < len(*right) {
		(*dst)[k] = (*right)[j]
		j++
		k++
	}
}
