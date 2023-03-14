package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkRecursiveInsertionSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go recursiveInsertionSort(arreglos[i], wg)
	}
}

func recursiveInsertionSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer tiempo.MedirTiempo(modelos.RECURSIVE_INSERTION_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.RecursiveInsertionSort(&arreglo.Arr, len(arreglo.Arr))
}
