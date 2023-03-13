package cmd

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/util"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

var archivo10Mil *os.File
var archivo50Mil *os.File
var archivo100Mil *os.File
var archivo500Mil *os.File
var archivo1Millon *os.File

var arreglo10Mil modelos.Arreglo
var arreglo50Mil modelos.Arreglo
var arreglo100Mil modelos.Arreglo
var arreglo500mil modelos.Arreglo
var arreglo1Millon modelos.Arreglo

var crearArreglosCmd = &cobra.Command{
	Use:   "crear-arreglos",
	Short: "Crear arreglos para los algoritmos",
	Long: `Crea arreglos de 1, 10 y 100 millones de
	elementos para ser utilizados en los diferentes
	algoritmos`,
	Run: crearArreglos,
}

func init() {
	rootCmd.AddCommand(crearArreglosCmd)
}

func crearArreglos(cmd *cobra.Command, args []string) {
	archivo10Mil = util.CrearArchivo("arreglo-10-mil.json")
	archivo50Mil = util.CrearArchivo("arreglo-50-mil.json")
	archivo100Mil = util.CrearArchivo("arreglo-100-mil.json")
	archivo500Mil = util.CrearArchivo("arreglo-500-mil.json")
	archivo1Millon = util.CrearArchivo("arreglo-1-millon.json")

	arreglo10Mil.Arr = util.GenerarArreglo(10000)

	arreglo50Mil.Arr = util.GenerarArreglo(50000)

	arreglo100Mil.Arr = util.GenerarArreglo(100000)

	arreglo500mil.Arr = util.GenerarArreglo(500000)

	arreglo1Millon.Arr = util.GenerarArreglo(1000000)

	encoder10Mil := json.NewEncoder(archivo10Mil)
	encoder50Mil := json.NewEncoder(archivo50Mil)
	encoder100Mil := json.NewEncoder(archivo100Mil)
	encoder500Mil := json.NewEncoder(archivo500Mil)
	encoder1Millon := json.NewEncoder(archivo1Millon)

	encoder10Mil.Encode(arreglo10Mil)
	encoder50Mil.Encode(arreglo50Mil)
	encoder100Mil.Encode(arreglo100Mil)
	encoder500Mil.Encode(arreglo500mil)
	encoder1Millon.Encode(arreglo1Millon)
}
