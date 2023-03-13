package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkQuickSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go quickSort(arreglos[i], wg)
	}
}

func quickSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Quick Sort 10 mil elementos"
	case 50000:
		titulo = "Quick Sort 50 mil elementos"
	case 100000:
		titulo = "Quick Sort 100 mil elementos"
	case 500000:
		titulo = "Quick Sort 500 mil elementos"
	case 1000000:
		titulo = "Quick Sort 1 millón de elementos"
	case 10000000:
		titulo = "Quick Sort 10 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()
	defer wg.Done()

	ordenamiento.QuickSort(&arreglo.Arr, 0, len(arreglo.Arr)-1)

}
