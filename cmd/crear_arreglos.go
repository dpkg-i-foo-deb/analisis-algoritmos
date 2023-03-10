package cmd

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/util"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

var archivo10Mil *os.File
var archivo100Mil *os.File
var archivo1Millon *os.File

var arreglo10Mil modelos.Arreglo
var arreglo100Mil modelos.Arreglo
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
	archivo100Mil = util.CrearArchivo("arreglo-100-mil.json")
	archivo1Millon = util.CrearArchivo("arreglo-1-millon.json")

	arreglo10Mil.Arr = util.GenerarArreglo(10000)

	arreglo100Mil.Arr = util.GenerarArreglo(100000)

	arreglo1Millon.Arr = util.GenerarArreglo(1000000)

	encoder10Mil := json.NewEncoder(archivo10Mil)
	encoder100Mil := json.NewEncoder(archivo100Mil)
	encoder1Millon := json.NewEncoder(archivo1Millon)

	encoder10Mil.Encode(arreglo10Mil)
	encoder100Mil.Encode(arreglo100Mil)
	encoder1Millon.Encode(arreglo1Millon)
}
