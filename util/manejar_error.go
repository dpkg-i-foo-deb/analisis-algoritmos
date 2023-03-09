package util

import "log"

func VerificarErrorDetener(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
