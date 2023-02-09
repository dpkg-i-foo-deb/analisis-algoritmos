package matrizsimetrica

// Determinar si una matriz es simétrica de una manera decente
func EsSimetricaDecente(matriz [][]int) bool {

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			if matriz[i][j] != matriz[j][i] {
				return false
			}
		}
	}

	return true
}

func EsSimetricaPeor(matriz [][]int) bool {
	res := true

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[i]); j++ {
			if matriz[i][j] != matriz[j][i] {
				res = false
			}
		}
	}

	return res
}
