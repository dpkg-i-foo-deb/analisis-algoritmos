package ordenamiento

// Creado por ChatGPT
// Adaptado para usar punteros
func BitonicSort(arr *[]int, low, cnt int, up bool) {
	if cnt > 1 {
		k := cnt / 2
		BitonicSort(arr, low, k, true)
		BitonicSort(arr, low+k, cnt-k, false)
		bitonicMerge(arr, low, cnt, up)
	}
}

func bitonicMerge(arr *[]int, low, cnt int, up bool) {
	if cnt > 1 {
		k := greatestPowerOfTwoLessThan(cnt)
		for i := low; i < low+cnt-k; i++ {
			if (up && (*arr)[i] > (*arr)[i+k]) || (!up && (*arr)[i] < (*arr)[i+k]) {
				swap(arr, i, i+k)
			}
		}
		bitonicMerge(arr, low, k, up)
		bitonicMerge(arr, low+k, cnt-k, up)
	}
}

func greatestPowerOfTwoLessThan(n int) int {
	k := 1
	for k < n {
		k = k << 1
	}
	return k >> 1
}

func swap(arr *[]int, i, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}
