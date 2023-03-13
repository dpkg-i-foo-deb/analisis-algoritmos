package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkInsertionSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go insertionSort(arreglos[i], wg)
	}
}

func insertionSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Insertion Sort 10 mil elementos"
	case 50000:
		titulo = "Insertion Sort 50 mil elementos"
	case 100000:
		titulo = "Insertion Sort 100 mil elementos"
	case 500000:
		titulo = "Insertion Sort 500 mil elementos"
	case 1000000:
		titulo = "Insertion Sort 1 millón de elementos"
	case 10000000:
		titulo = "Insertion Sort 10 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo, modelos.INSERTION_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.InsertionSort(&arreglo.Arr)
}
