package ordenamiento

// Creado por ChatGPT
// Adaptado para usar punteros
func GnomeSort(arr *[]int) {
	i := new(int)
	*i = 1
	j := new(int)
	*j = 2
	for *i < len(*arr) {
		if (*arr)[*i-1] <= (*arr)[*i] {
			*i = *j
			*j++
		} else {
			(*arr)[*i-1], (*arr)[*i] = (*arr)[*i], (*arr)[*i-1]
			*i--
			if *i == 0 {
				*i = *j
				*j++
			}
		}
	}
}
