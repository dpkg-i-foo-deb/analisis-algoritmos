package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkStrandSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go strandSort(arreglos[i], wg)
	}
}

func strandSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer tiempo.MedirTiempo(modelos.STRAND_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.StrandSort(&arreglo.Arr)
}
