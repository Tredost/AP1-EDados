package loja

import (
	"net/http"
	// IMPORTAR PROCESSAMENTO
)

func HandleAbrirLoja(w http.ResponseWriter, r *http.Request) {
	AbrirLoja(w, r)
}

func HandleFecharLoja(w http.ResponseWriter, r *http.Request) {
	FecharLoja(w, r)
}
