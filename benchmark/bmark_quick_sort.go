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

	defer tiempo.MedirTiempo(modelos.QUICK_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.QuickSort(&arreglo.Arr, 0, len(arreglo.Arr)-1)

}
