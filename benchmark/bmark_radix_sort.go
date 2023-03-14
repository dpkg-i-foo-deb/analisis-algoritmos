package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkRadixSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go radixSort(arreglos[i], wg)
	}
}

func radixSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer tiempo.MedirTiempo(modelos.RADIX_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.RadixSort(&arreglo.Arr)
}
