package matrices

func RestaNormalita(a [][]int, b [][]int) [][]int {
	resultado := make([][]int, 0)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			valor := make([]int, 0)

			valor = append(valor, a[i][j]-b[i][j])

			resultado = append(resultado, valor)

		}
	}

	return resultado
}
