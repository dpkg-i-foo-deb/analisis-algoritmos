package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkBitonicSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go bitonicSort(arreglos[i], wg)
	}
}

func bitonicSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {
	defer tiempo.MedirTiempo(modelos.BITONIC_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.BitonicSort(&arreglo.Arr, 0, len(arreglo.Arr), true)
}
