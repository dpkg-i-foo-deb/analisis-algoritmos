package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkSelectionSort(arreglos []modelos.Arreglo) {
	for i := range arreglos {
		go selectionSort(arreglos[i])
	}
}

func selectionSort(arreglo modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Selection Sort 10 mil elementos"
	case 100000:
		titulo = "Selection Sort 100 mil elementos"
	case 1000000:
		titulo = "Selection Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.SelectionSort(&arreglo.Arr)
}
