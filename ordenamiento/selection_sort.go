package ordenamiento

// Generado por ChatGPT
// Adaptado para utilizar punteros
func SelectionSort(arr *[]int) {
	n := len(*arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if (*arr)[j] < (*arr)[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			// swap arr[i] and arr[minIndex]
			(*arr)[i], (*arr)[minIndex] = (*arr)[minIndex], (*arr)[i]
		}
	}
}
