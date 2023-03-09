package ordenamiento

// Creado por ChatGPT
// Adaptado para usar punteros
func BinaryInsertionSort(arr *[]int) {
	for i := 1; i < len(*arr); i++ {
		x := (*arr)[i]
		left := new(int)
		*left = 0
		right := new(int)
		*right = i - 1
		for *left <= *right {
			mid := (*left + *right) / 2
			if x < (*arr)[mid] {
				*right = mid - 1
			} else {
				*left = mid + 1
			}
		}
		for j := i - 1; j >= *left; j-- {
			(*arr)[j+1] = (*arr)[j]
		}
		(*arr)[*left] = x
	}
}
