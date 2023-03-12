package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkStrandSort(arreglos []modelos.Arreglo) {
	for i := range arreglos {
		strandSort(arreglos[i])
	}
}

func strandSort(arreglo modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Strand Sort 10 mil elementos"
	case 100000:
		titulo = "Strand Sort 100 mil elementos"
	case 1000000:
		titulo = "Strand Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.StrandSort(&arreglo.Arr)
}
