package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkMergeSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go mergeSort(arreglos[i], wg)
	}
}

func mergeSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer tiempo.MedirTiempo(modelos.MERGE_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.MergeSort(&arreglo.Arr)
}
