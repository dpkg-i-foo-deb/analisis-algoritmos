package pruebas

import (
	"analisis-algoritmos/matrices"
	"analisis-algoritmos/tiempo"
)

func RestaMatrices() {

	a := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	b := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	probarRestaNormal(a, b)

}

func probarRestaNormal(a [][]int, b [][]int) {

	defer tiempo.MedirTiempo("Resta de matrices")()

	matrices.RestaNormalita(a, b)
}
