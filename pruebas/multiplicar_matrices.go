package pruebas

import (
	"analisis-algoritmos/matrices"
	"analisis-algoritmos/tiempo"
)

func MultiplicarMatricesPrueba() {
	a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	b := [][]int{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	multiplicarPeor(a, b)
}

func multiplicarPeor(a [][]int, b [][]int) {

	defer tiempo.MedirTiempo("Peor multiplicación de matrices")()

	matrices.MultiplicarMatrices(a, b)
}
