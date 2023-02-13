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

	aFloat := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	bFloat := [][]float64{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	multiplicarPeor(a, b)
	MultiplicarMejor(aFloat, bFloat)
}

func multiplicarPeor(a [][]int, b [][]int) {

	defer tiempo.MedirTiempo("Peor multiplicación de matrices")()

	matrices.MultiplicarMatrices(a, b)
}

func MultiplicarMejor(a [][]float64, b [][]float64) {
	defer tiempo.MedirTiempo("Multiplicar matrices un tris mejor")()

	matrices.MultiplicarKahan(a, b)
}
