package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkStoogeSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go stoogeSort(arreglos[i], wg)
	}
}

func stoogeSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer wg.Done()

	if len(arreglo.Arr) > 10002 {
		return
	}
	defer tiempo.MedirTiempo(modelos.STOOGE_SORT, len(arreglo.Arr))()

	ordenamiento.StoogeSort(&arreglo.Arr, 0, len(arreglo.Arr)-1)
}
