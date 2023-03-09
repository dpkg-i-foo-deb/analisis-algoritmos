package matrices

func SumitaNormal(a [][]int, b [][]int) [][]int {

	resultado := make([][]int, 0)

	for i := 0; i < len(a); i++ {
		valor := make([]int, 0)
		for j := 0; j < len(a[0]); j++ {
			valor = append(valor, a[i][j]+b[i][j])
		}
		resultado = append(resultado, valor)
	}

	return resultado
}
