package pruebas

import (
	"analisis-algoritmos/matrices"
	"analisis-algoritmos/tiempo"
)

func Multiplicar() {

	matriz := [][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}

	escalar := 8

	multiplicacionNormalita(matriz, escalar)

}

func multiplicacionNormalita(matriz [][]int, escalar int) {

	defer tiempo.MedirTiempo("Multiplicar por escalar")()

	matrices.MultiplicarEscalar(matriz, escalar)

}
