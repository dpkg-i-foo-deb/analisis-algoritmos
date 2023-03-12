package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkRecursiveInsertionSort(arreglos []modelos.Arreglo) {
	for i := range arreglos {
		go recursiveInsertionSort(arreglos[i])
	}
}

func recursiveInsertionSort(arreglo modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Recursive Insertion Sort 10 mil elementos"
	case 100000:
		titulo = "Recursive Insertion Sort 100 mil elementos"
	case 1000000:
		titulo = "Recursive Insertion Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.RecursiveInsertionSort(&arreglo.Arr, len(arreglo.Arr))
}
