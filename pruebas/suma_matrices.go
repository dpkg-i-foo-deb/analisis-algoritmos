package pruebas

import (
	"analisis-algoritmos/matrices"
	"analisis-algoritmos/tiempo"
)

func SumaMatrices() {
	a := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	b := [][]int{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}

	probarSumitaNormal(a, b)
}

func probarSumitaNormal(a [][]int, b [][]int) {
	defer tiempo.MedirTiempo("Suma de matrices normalita")()
	matrices.SumitaNormal(a, b)
}
