package modelos

import "time"

type Resultado struct {
	Titulo   string        `json:"nombre"`
	Duracion time.Duration `json:"duracion"`
}
