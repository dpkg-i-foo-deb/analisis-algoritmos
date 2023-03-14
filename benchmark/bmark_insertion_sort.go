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

	defer tiempo.MedirTiempo(modelos.INSERTION_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.InsertionSort(&arreglo.Arr)
}
