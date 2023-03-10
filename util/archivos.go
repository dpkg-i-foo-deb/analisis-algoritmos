package util

import "os"

func CrearArchivo(nombre string) *os.File {
	archivo, err := os.Create(nombre)

	VerificarErrorDetener(err)

	return archivo
}

func AbrirArchivo(nombre string) *os.File {
	archivo, err := os.Open(nombre)

	VerificarErrorDetener(err)

	return archivo
}
