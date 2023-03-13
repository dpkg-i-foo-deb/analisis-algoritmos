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
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Shell Sort 10 mil elementos"
	case 50000:
		titulo = "Shell Sort 50 mil elementos"
	case 100000:
		titulo = "Shell Sort 100 mil elementos"
	case 500000:
		titulo = "Shell Sort 500 mil elementos"
	case 1000000:
		titulo = "Shell Sort 1 millón de elementos"
	case 10000000:
		titulo = "Shell Sort 10 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo, modelos.SHELL_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.ShellSort(&arreglo.Arr)
}
