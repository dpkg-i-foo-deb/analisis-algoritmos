package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkQuickSort(arreglos []*modelos.Arreglo) {
	for i := range arreglos {
		quickSort(arreglos[i])
	}
}

func quickSort(arreglo *modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 1000000:
		titulo = "Quick Sort 1 millón de elementos"
	case 10000000:
		titulo = "Quick Sort 10 millones de elementos"
	case 100000000:
		titulo = "Quick Sort 100 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.QuickSort(&arreglo.Arr, 0, len(arreglo.Arr)-1)

}
