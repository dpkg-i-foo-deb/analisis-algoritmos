package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkInsertionSort(arreglos []*modelos.Arreglo) {
	for i := range arreglos {
		insertionSort(arreglos[i])
	}
}

func insertionSort(arreglo *modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 1000000:
		titulo = "Insertion Sort 1 millón de elementos"
	case 10000000:
		titulo = "Insertion Sort 10 millones de elementos"
	case 100000000:
		titulo = "Insertion Sort 100 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.InsertionSort(&arreglo.Arr)
}
