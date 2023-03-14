package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkHeapSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go heapSort(arreglos[i], wg)
	}
}

func heapSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer tiempo.MedirTiempo(modelos.HEAP_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.HeapSort(&arreglo.Arr)
}
