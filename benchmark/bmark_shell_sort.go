package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkShellSort(arreglos []*modelos.Arreglo) {
	for i := range arreglos {
		shellSort(arreglos[i])
	}
}

func shellSort(arreglo *modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 1000000:
		titulo = "Shell Sort 1 millón de elementos"
	case 10000000:
		titulo = "Shell Sort 10 millones de elementos"
	case 100000000:
		titulo = "Shell Sort 100 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.ShellSort(&arreglo.Arr)
}
