package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkShellSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go shellSort(arreglos[i], wg)
	}
}

func shellSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer tiempo.MedirTiempo(modelos.SHELL_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.ShellSort(&arreglo.Arr)
}
