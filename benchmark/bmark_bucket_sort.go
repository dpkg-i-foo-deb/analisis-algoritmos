package benchmark

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/ordenamiento"
	"analisis-algoritmos/tiempo"
	"sync"
)

func BmarkBucketSort(arreglos []modelos.Arreglo, wg *sync.WaitGroup) {
	for i := range arreglos {
		wg.Add(1)
		go bucketSort(arreglos[i], wg)
	}
}

func bucketSort(arreglo modelos.Arreglo, wg *sync.WaitGroup) {

	defer tiempo.MedirTiempo(modelos.BUCKET_SORT, len(arreglo.Arr))()
	defer wg.Done()

	ordenamiento.BucketSort(&arreglo.Arr)
}
