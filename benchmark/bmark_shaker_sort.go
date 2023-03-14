package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkShakerSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go shakerSort(arreglos[i], wg)
	}
}

func shakerSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer tiempo.MedirTiempo(modelos.SHAKER_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.ShakerSort(&arreglo.Arr)
}
