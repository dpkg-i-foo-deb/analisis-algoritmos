package pruebas

import (
	"analisis-algoritmos/matrices"
	"analisis-algoritmos/tiempo"
)

func MatrizSimetrica() {
	m := [][]int{
		{1, 6, 0},
		{6, -3, 8},
		{0, 1, 4},
	}

	verificarMejorSimetrica(&m)
	verificarPeorSimetrica(&m)

}

func verificarMejorSimetrica(matriz *[][]int) {
	defer tiempo.MedirTiempo("Mejor matriz simétrica")()

	matrices.EsSimetricaDecente(*matriz)
}

func verificarPeorSimetrica(matriz *[][]int) {
	defer tiempo.MedirTiempo("Peor matriz simétrica")()

	matrices.EsSimetricaPeor(*matriz)
}
