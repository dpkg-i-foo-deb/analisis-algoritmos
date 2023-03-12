package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
)

func BmarkBucketSort(arreglos []modelos.Arreglo) {
	for i := range arreglos {
		go bucketSort(arreglos[i])
	}
}

func bucketSort(arreglo modelos.Arreglo) {
	var titulo string

	switch len(arreglo.Arr) {
	case 10000:
		titulo = "Bucket Sort 10 mil elementos"
	case 100000:
		titulo = "Bucket Sort 100 mil elementos"
	case 1000000:
		titulo = "Bucket Sort 1 millón de elementos"
	}

	defer tiempo.MedirTiempo(titulo)()

	ordenamiento.BucketSort(&arreglo.Arr)
}
