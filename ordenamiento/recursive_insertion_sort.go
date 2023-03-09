package ordenamiento

// Generado por ChatGPT
// Adaptado para usar punteros
func RecursiveInsertionSort(arr *[]int, n int) {
	if n <= 1 {
		return
	}

	RecursiveInsertionSort(arr, n-1)

	last := (*arr)[n-1]
	j := n - 2

	for j >= 0 && (*arr)[j] > last {
		(*arr)[j+1] = (*arr)[j]
		j--
	}

	(*arr)[j+1] = last
}
