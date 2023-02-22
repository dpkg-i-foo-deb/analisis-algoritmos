package pruebas

import (
	matrizespiral "analisis-algoritmos/matriz_espiral"
	"analisis-algoritmos/tiempo"
	"analisis-algoritmos/util"
)

func Espiral() {
	defer tiempo.MedirTiempo("Llenar la matriz en espiral")()

	matriz := matrizespiral.LlenarEspiral(5)

	util.ImprimirMatriz(matriz)
}
