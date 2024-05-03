package IANZINHO

import (
	p "IANZINHO/processamento"
	"net/http"
)

func HandleIncluirPedido(w http.ResponseWriter, r *http.Request) {
	p.IncluirPedido(w, r)
}

func HandleObterPedidosAtivos(w http.ResponseWriter, r *http.Request) {
	p.ObterPedidosAtivos(w, r)
}
