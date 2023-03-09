package ordenamiento

// Creado por ChatGPT
// Adaptado para usar punteros
func QuickSort(arr *[]int, low, high int) {
	if low < high {
		// Partition the array
		pivotIndex := partition(arr, low, high)

		// Recursively sort the left and right sub-arrays
		QuickSort(arr, low, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr *[]int, low, high int) int {
	// Choose the rightmost element as the pivot
	pivot := (*arr)[high]

	// Initialize the pivot index to the leftmost index
	pivotIndex := low

	// Partition the array
	for i := low; i < high; i++ {
		if (*arr)[i] < pivot {
			// Swap arr[i] and arr[pivotIndex]
			temp := (*arr)[i]
			(*arr)[i] = (*arr)[pivotIndex]
			(*arr)[pivotIndex] = temp

			// Increment pivotIndex
			pivotIndex++
		}
	}

	// Swap arr[high] and arr[pivotIndex]
	temp := (*arr)[high]
	(*arr)[high] = (*arr)[pivotIndex]
	(*arr)[pivotIndex] = temp

	// Return the pivot index
	return pivotIndex
}
