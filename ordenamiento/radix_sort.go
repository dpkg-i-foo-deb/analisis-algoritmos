package ordenamiento

// Generado por ChatGPT
// Adaptado para usar punteros
func RadixSort(arr *[]int) {
	if len(*arr) == 0 {
		return
	}

	maxVal := (*arr)[0]
	for i := 1; i < len(*arr); i++ {
		if (*arr)[i] > maxVal {
			maxVal = (*arr)[i]
		}
	}

	exp := 1
	for maxVal/exp > 0 {
		countArr := make([]int, 10)

		for i := 0; i < len(*arr); i++ {
			countArr[((*arr)[i]/exp)%10]++
		}

		for i := 1; i < 10; i++ {
			countArr[i] += countArr[i-1]
		}

		output := make([]int, len(*arr))
		for i := len(*arr) - 1; i >= 0; i-- {
			output[countArr[((*arr)[i]/exp)%10]-1] = (*arr)[i]
			countArr[((*arr)[i]/exp)%10]--
		}

		for i := 0; i < len(*arr); i++ {
			(*arr)[i] = output[i]
		}

		exp *= 10
	}
}
