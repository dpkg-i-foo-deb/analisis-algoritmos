package cmd

import (
	"analisis-algoritmos/modelos"
	"analisis-algoritmos/util"
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

var archivo1Millon *os.File
var archivo10Millones *os.File
var archivo100Millones *os.File

var arreglo1Millon modelos.Arreglo
var arreglo10Millones modelos.Arreglo
var arreglo100Millones modelos.Arreglo

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
	archivo1Millon = util.CrearArchivo("arreglo-1-millon.json")
	archivo10Millones = util.CrearArchivo("arreglo-10-millones.json")
	archivo100Millones = util.CrearArchivo("arreglo-100-millones.json")

	arreglo1Millon.Arr = util.GenerarArreglo(4, 1000000)

	arreglo10Millones.Arr = util.GenerarArreglo(4, 10000000)

	arreglo100Millones.Arr = util.GenerarArreglo(4, 100000000)

	encoder1Millon := json.NewEncoder(archivo1Millon)
	encoder10Millones := json.NewEncoder(archivo10Millones)
	encoder100Millones := json.NewEncoder(archivo100Millones)

	encoder1Millon.Encode(arreglo1Millon)
	encoder10Millones.Encode(arreglo10Millones)
	encoder100Millones.Encode(arreglo100Millones)
}
