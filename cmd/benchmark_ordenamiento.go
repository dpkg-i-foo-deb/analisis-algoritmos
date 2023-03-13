package cmd

import (
	"analisis-algoritmos/benchmark"
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/resultados"
	"analisis-algoritmos/util"
	"encoding/json"
	"log"
	"sync"

	"github.com/spf13/cobra"
)

var benchmarkOrdenamientoCmd = &cobra.Command{
	Use:   "benchmark-ordenamiento",
	Short: "Comparar tiempos de ordenamiento",
	Long: `Compara el tiempo de ejecución de diferentes algoritmos
	de ordenamiento utilizando arreglos de 1, 10 y 100 millones de elementos
	(Deben haber sido creados previamente)`,
	Run: benchmarkOrdenamiento,
}

var benchmarks []func([]modelos.Arreglo, *sync.WaitGroup)

var arreglosBenchmark []modelos.Arreglo

func init() {
	rootCmd.AddCommand(benchmarkOrdenamientoCmd)
	benchmarks = append(benchmarks, benchmark.BmarkInsertionSort)
	benchmarks = append(benchmarks, benchmark.BmarkBinaryInsertionSort)
	benchmarks = append(benchmarks, benchmark.BmarkBitonicSort)
	benchmarks = append(benchmarks, benchmark.BmarkBubbleSort)
	benchmarks = append(benchmarks, benchmark.BmarkBucketSort)
	benchmarks = append(benchmarks, benchmark.BmarkGnomeSort)
	benchmarks = append(benchmarks, benchmark.BmarkHeapSort)
	benchmarks = append(benchmarks, benchmark.BmarkQuickSort)
	benchmarks = append(benchmarks, benchmark.BmarkRadixSort)
	benchmarks = append(benchmarks, benchmark.BmarkRecursiveInsertionSort)
	benchmarks = append(benchmarks, benchmark.BmarkSelectionSort)
	benchmarks = append(benchmarks, benchmark.BmarkShakerSort)
	benchmarks = append(benchmarks, benchmark.BmarkShellSort)
	benchmarks = append(benchmarks, benchmark.BmarkStoogeSort)
	benchmarks = append(benchmarks, benchmark.BmarkStrandSort)
	benchmarks = append(benchmarks, benchmark.BmarkMergeSort)

}

func benchmarkOrdenamiento(cmd *cobra.Command, args []string) {
	abrirArchivos()
	leerAreglos()
	ejecutarBenchmarks()
	resultados.Consolidar()
	resultados.EscribirResultado()
}

func abrirArchivos() {
	log.Println("Buscando los archivos...")
	archivo10Mil = util.AbrirArchivo("arreglo-10-mil.json")
	archivo50Mil = util.AbrirArchivo("arreglo-50-mil.json")
	archivo100Mil = util.AbrirArchivo("arreglo-100-mil.json")
	archivo500Mil = util.AbrirArchivo("arreglo-500-mil.json")
	archivo1Millon = util.AbrirArchivo("arreglo-1-millon.json")
	archivo10Millones = util.AbrirArchivo("arreglo-10-millones.json")

}

func leerAreglos() {
	log.Println("Leyendo los arreglos...")
	decoder10Mil := json.NewDecoder(archivo10Mil)
	decoder50Mil := json.NewDecoder(archivo50Mil)
	decoder100Mil := json.NewDecoder(archivo100Mil)
	decoder500Mil := json.NewDecoder(archivo500Mil)
	decoder1Millon := json.NewDecoder(archivo1Millon)
	decoder10Millones := json.NewDecoder(archivo10Millones)

	log.Println("10 mil elementos...")
	err := decoder10Mil.Decode(&arreglo10Mil)

	util.VerificarErrorDetener(err)

	log.Println("50 mil elementos...")
	err = decoder50Mil.Decode(&arreglo50Mil)

	util.VerificarErrorDetener(err)

	log.Println("100 mil elementos...")
	err = decoder100Mil.Decode(&arreglo100Mil)

	util.VerificarErrorDetener(err)

	log.Println("500 mil elementos...")
	err = decoder500Mil.Decode(&arreglo500mil)

	util.VerificarErrorDetener(err)

	log.Println("1 millón de elementos...")
	err = decoder1Millon.Decode(&arreglo1Millon)

	util.VerificarErrorDetener(err)

	log.Println("10 millones de elementos...")
	err = decoder10Millones.Decode(&arreglo10Millones)

	util.VerificarErrorDetener(err)

	arreglosBenchmark = append(arreglosBenchmark, arreglo10Mil, arreglo50Mil, arreglo100Mil, arreglo500mil, arreglo1Millon, arreglo10Millones)
}

func ejecutarBenchmarks() {
	log.Println("Ejecutando los benchmark, tomará un rato...")

	var wg sync.WaitGroup

	for i := range benchmarks {

		copia := make([]modelos.Arreglo, len(arreglosBenchmark), cap(arreglosBenchmark))

		copy(copia, arreglosBenchmark)

		benchmarks[i](copia, &wg)

	}

	wg.Wait()
}
