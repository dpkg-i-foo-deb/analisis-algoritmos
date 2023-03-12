package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkBitonicSort(arreglos []modelos.Arreglo) {
	for i := range arreglos {
		go bitonicSort(arreglos[i])
	}
}

func bitonicSort(arreglo modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Bitonic Sort 10 mil elementos"
	case 100000:
		titulo = "Bitonic Sort 100 mil elementos"
	case 1000000:
		titulo = "Bitonic Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.BitonicSort(&arreglo.Arr, 0, len(arreglo.Arr), true)
}
