package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkQuickSort(arreglos []modelos.Arreglo) {
	for i := range arreglos {
		quickSort(arreglos[i])
	}
}

func quickSort(arreglo modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Quick Sort 10 mil elementos"
	case 100000:
		titulo = "Quick Sort 100 mil elementos"
	case 1000000:
		titulo = "Quick Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.QuickSort(&arreglo.Arr, 0, len(arreglo.Arr)-1)

}
