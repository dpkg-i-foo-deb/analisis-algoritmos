package pruebas

import (
	matrizsimetrica "analisis-algoritmos/matriz_simetrica"
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

	matrizsimetrica.EsSimetricaDecente(*matriz)
}

func verificarPeorSimetrica(matriz *[][]int) {
	defer tiempo.MedirTiempo("Peor matriz simétrica")()

	matrizsimetrica.EsSimetricaPeor(*matriz)
}
