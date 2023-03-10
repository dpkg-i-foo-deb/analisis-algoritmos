package ordenamiento

// Creado por ChatGPT
// Adaptado para usar punteros
func BitonicSort(arr *[]int, low, cnt int, up bool) {
	if cnt <= 1 {
		return
	}

	m := cnt / 2
	BitonicSort(arr, low, m, !up)
	BitonicSort(arr, low+m, cnt-m, up)
	BitonicMerge(arr, low, cnt, up)
}

// BitonicMerge merges two bitonic sequences in non-decreasing or non-increasing order
func BitonicMerge(arr *[]int, low, cnt int, up bool) {
	if cnt <= 1 {
		return
	}

	m := cnt / 2
	for i := low; i < low+m; i++ {
		if (*arr)[i] > (*arr)[i+m] == up {
			(*arr)[i], (*arr)[i+m] = (*arr)[i+m], (*arr)[i]
		}
	}
	BitonicMerge(arr, low, m, up)
	BitonicMerge(arr, low+m, cnt-m, up)
}
