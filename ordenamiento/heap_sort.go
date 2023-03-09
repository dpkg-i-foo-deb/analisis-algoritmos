package ordenamiento

//Creado por ChatGPT
//Adaptado para usar punteros

func HeapSort(arr *[]int) {
	// Build a max heap
	buildMaxHeap(arr)

	// Sort the array
	for i := len(*arr) - 1; i > 0; i-- {
		// Swap the root element with the last element
		temp := (*arr)[i]
		(*arr)[i] = (*arr)[0]
		(*arr)[0] = temp

		// Restore the heap property
		heapify(arr, 0, i)
	}
}

func buildMaxHeap(arr *[]int) {
	for i := len(*arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i, len(*arr))
	}
}

func heapify(arr *[]int, i, size int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	// Find the largest element among i, left, and right
	if left < size && (*arr)[left] > (*arr)[largest] {
		largest = left
	}
	if right < size && (*arr)[right] > (*arr)[largest] {
		largest = right
	}

	// If i is not the largest, swap it with the largest element and continue heapifying
	if largest != i {
		temp := (*arr)[i]
		(*arr)[i] = (*arr)[largest]
		(*arr)[largest] = temp
		heapify(arr, largest, size)
	}
}
