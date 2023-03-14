package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkGnomeSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go gnomeSort(arreglos[i], wg)
	}
}

func gnomeSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer tiempo.MedirTiempo(modelos.GNOME_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.GnomeSort(&arreglo.Arr)
}
