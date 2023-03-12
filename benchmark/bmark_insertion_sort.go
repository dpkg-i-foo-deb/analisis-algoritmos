package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkInsertionSort(arreglos []modelos.Arreglo) {
	for i := range arreglos {
		go insertionSort(arreglos[i])
	}
}

func insertionSort(arreglo modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Insertion Sort 10 mil elementos"
	case 100000:
		titulo = "Insertion Sort 100 mil elementos"
	case 1000000:
		titulo = "Insertion Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.InsertionSort(&arreglo.Arr)
}
