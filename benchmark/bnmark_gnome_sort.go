package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkGnomeSort(arreglos []*modelos.Arreglo) {
	for i := range arreglos {
		gnomeSort(arreglos[i])
	}
}

func gnomeSort(arreglo *modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 1000000:
		titulo = "Gnome Sort 1 millón de elementos"
	case 10000000:
		titulo = "Gnome Sort 10 millones de elementos"
	case 100000000:
		titulo = "Gnome Sort 100 millones de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.GnomeSort(&arreglo.Arr)
}