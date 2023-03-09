package matrices

func MultiplicarMatrices(a [][]int, b [][]int) [][]int {

	resultado := make([][]int, 0)

	for i := 0; i < len(a); i++ {
		valor := make([]int, 0)
		for j := 0; j < len(b[0]); j++ {
			productoria := 0
			for k := 0; k < len(b); k++ {
				productoria += a[i][k] * b[k][j]
			}
			valor = append(valor, productoria)
		}
		resultado = append(resultado, valor)
	}
	return resultado
}

func MultiplicarKahan(a [][]float64, b [][]float64) [][]float64 {
	resultado := make([][]float64, 0)
	correccion := 0.0
	for i := 0; i < len(a); i++ {
		valor := make([]float64, 0)
		for j := 0; j < len(b[0]); j++ {

			suma := 0.0
			err := 0.0

			for k := 0; k < len(b); k++ {
				err = err + a[i][k]*b[k][j]
				correccion = suma + err
				err = (suma - correccion) + err
				suma = correccion
			}

			valor = append(valor, suma)

		}

		resultado = append(resultado, valor)
	}

	return resultado
}
