package ordenamiento

//Creado por ChatGPT
//Adaptado para usar punteros

func StoogeSort(arr *[]int, left, right int) {
	if (*arr)[left] > (*arr)[right] {
		// Swap the left and right elements
		temp := (*arr)[left]
		(*arr)[left] = (*arr)[right]
		(*arr)[right] = temp
	}

	if right-left > 1 {
		// Sort the first two-thirds of the array
		t := (right - left + 1) / 3
		StoogeSort(arr, left, right-t)

		// Sort the last two-thirds of the array
		StoogeSort(arr, left+t, right)

		// Sort the first two-thirds of the array again
		StoogeSort(arr, left, right-t)
	}
}
