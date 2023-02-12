package matrices

func MultiplicarEscalar(matriz [][]int, escalar int) [][]int {
	resultado := make([][]int, 0)

	for i := 0; i < len(matriz); i++ {
		valor := make([]int, 0)
		for j := 0; j < len(matriz[0]); j++ {
			valor = append(valor, matriz[i][j]*escalar)
		}
		resultado = append(resultado, valor)
	}

	return resultado
}
