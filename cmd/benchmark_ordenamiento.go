package cmd

import (
	"analisis-algoritmos/benchmark"
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/util"
	"encoding/json"
	"log"

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

var arreglosBenchmark []*modelos.Arreglo

func init() {
	rootCmd.AddCommand(benchmarkOrdenamientoCmd)
}

func benchmarkOrdenamiento(cmd *cobra.Command, args []string) {
	abrirArchivos()
	leerAreglos()
	log.Println("Ejecutando los benchmark, tomará un rato...")
	benchmark.BmarkBinaryInsertionSort(arreglosBenchmark)
	benchmark.BmarkBitonicSort(arreglosBenchmark)
	benchmark.BmarkBubbleSort(arreglosBenchmark)
	benchmark.BmarkBucketSort(arreglosBenchmark)
	benchmark.BmarkGnomeSort(arreglosBenchmark)
	benchmark.BmarkHeapSort(arreglosBenchmark)
	benchmark.BmarkBinaryInsertionSort(arreglosBenchmark)
	benchmark.BmarkQuickSort(arreglosBenchmark)
	benchmark.BmarkRadixSort(arreglosBenchmark)
	benchmark.BmarkRecursiveInsertionSort(arreglosBenchmark)
	benchmark.BmarkSelectionSort(arreglosBenchmark)
	benchmark.BmarkShakerSort(arreglosBenchmark)
	benchmark.BmarkShellSort(arreglosBenchmark)
	benchmark.BmarkStoogeSort(arreglosBenchmark)
	benchmark.BmarkStrandSort(arreglosBenchmark)
}

func abrirArchivos() {
	log.Println("Buscando los archivos...")
	archivo1Millon = util.AbrirArchivo("arreglo-1-millon.json")
	archivo10Millones = util.AbrirArchivo("arreglo-10-millones.json")
	archivo100Millones = util.AbrirArchivo("arreglo-100-millones.json")
}

func leerAreglos() {
	log.Println("Leyendo los arreglos...")
	decoder1Millon := json.NewDecoder(archivo1Millon)
	decoder10Millones := json.NewDecoder(archivo10Millones)
	decoder100Millones := json.NewDecoder(archivo100Millones)

	log.Println("1 millón de elementos...")
	err := decoder1Millon.Decode(&arreglo1Millon)

	util.VerificarErrorDetener(err)

	log.Println("10 millones de elementos...")
	err = decoder10Millones.Decode(&arreglo10Millones)

	util.VerificarErrorDetener(err)

	log.Println("100 millones de elementos...")
	err = decoder100Millones.Decode(&arreglo100Millones)

	util.VerificarErrorDetener(err)

	arreglosBenchmark = append(arreglosBenchmark, &arreglo1Millon, &arreglo10Millones, &arreglo100Millones)
}
