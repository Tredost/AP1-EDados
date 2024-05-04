package IANZINHO

import (
	p "IANZINHO/processamento"
	"net/http"
)

func HandleAbrirLoja(w http.ResponseWriter, r *http.Request) {
	p.AbrirLoja(w, r)
}

func HandleFecharLoja(w http.ResponseWriter, r *http.Request) {
	p.FecharLoja(w, r)
}
