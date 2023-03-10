package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkMergeSort(arreglos []*modelos.Arreglo) {
	for i := range arreglos {
		mergeSort(arreglos[i])
	}
}

func mergeSort(arreglo *modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Merge Sort 10 mil elementos"
	case 100000:
		titulo = "Merge Sort 100 mil elementos"
	case 1000000:
		titulo = "Merge Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.MergeSort(&arreglo.Arr)
}
