package matrizespiral

func LlenarEspiral(tam int) [][]int {

	matriz := make([][]int, tam)

	for i := range matriz {
		matriz[i] = make([]int, tam)
	}

	cuenta := 1
	limI := tam
	limJ := tam
	minI := 0
	minJ := 0
	abajo := true
	derecha := false
	arriba := false
	izquierda := false

	for i := 0; i < tam+2; {

		for j := 0; j < tam+2; {

			matriz[i][j] = cuenta
			cuenta++
			if cuenta == tam*tam+1 {
				return matriz
			}

			if abajo && i < limI-1 {
				i++
			}

			if abajo && i == limI-1 {
				derecha = true
				abajo = false
				limI--
				continue
			}

			if derecha && j < limJ-1 {
				j++
			}

			if derecha && j == limJ-1 {
				arriba = true
				derecha = false
				limJ--
				minJ++
				continue
			}

			if arriba && i > minI {
				i--
			}

			if arriba && i == minI {
				izquierda = true
				arriba = false
				minI++
				continue
			}

			if izquierda && j > minJ {
				j--
			}

			if izquierda && j == minJ {
				abajo = true
				izquierda = false
				continue
			}

		}
	}

	return matriz

}
